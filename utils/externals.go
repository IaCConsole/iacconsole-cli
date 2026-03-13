package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *State) GetStringFromViperByOrgOrDefault(keyName string) string {
	if viper.IsSet(s.OrgName + "." + keyName) {
		return viper.GetString(s.OrgName + "." + keyName)
	} else {
		return viper.GetString("defaults." + keyName)
	}
}

func (s *State) GetObjectFromViperByOrgOrDefault(keyName string) map[string]any {
	if viper.IsSet(s.OrgName + "." + keyName) {
		return viper.GetStringMap(s.OrgName + "." + keyName)
	} else {
		return viper.GetStringMap("defaults." + keyName)
	}
}

func (s *State) SetupBackendConfig() map[string]interface{} {
	var stateS3Path string
	if !viper.IsSet(s.OrgName + ".backend") {
		stateS3Path = stateS3Path + "org_" + s.OrgName + "/"
	}

	for _, dimension := range s.UnitManifest.Dimensions {
		stateS3Path = stateS3Path + dimension + "_" + s.ParsedDimensions[dimension] + "/"
	}
	s.StateS3Path = stateS3Path + s.UnitName + ".tfstate"

	backendConfig := s.GetObjectFromViperByOrgOrDefault("backend")
	if len(backendConfig) == 0 {
		log.Println("no backend config provied!")
	}

	var backendConfigMap = make(map[string]interface{}, len(backendConfig))
	for param, value := range backendConfig {
		backendConfigMap[param] = strings.Replace(value.(string), "$iacconsole_state_path", s.StateS3Path, 1)
	}

	return backendConfigMap
}

func (s *State) GetDimData(ctx context.Context, dimensionKey string, dimensionValue string, skipOnNotFound bool) (map[string]interface{}, error) {
	var dimensionJsonMap map[string]interface{}

	if s.IacconsoleApiUrl == "" {
		inventroyJsonPath := s.InventoryPath + "/" + dimensionKey + "/" + dimensionValue + ".json"
		dimensionJsonBytes, err := os.ReadFile(inventroyJsonPath)
		if err != nil {
			if os.IsNotExist(err) && skipOnNotFound {
				log.Println("inventory files: Optional dimension " + s.OrgName + "/" + dimensionKey + "/" + dimensionValue + " not found, skipping")
				return dimensionJsonMap, nil
			}
			return nil, err
		}
		err = json.Unmarshal(dimensionJsonBytes, &dimensionJsonMap)
		if err != nil {
			return nil, err
		}
	} else {
		req, err := http.NewRequestWithContext(ctx, "GET", s.IacconsoleApiUrl + "/v1/dimension/" + s.OrgName + "/" + dimensionKey + "/" + dimensionValue + "?workspace=" + s.Workspace + "&fallbacktomaster=true", nil)
		if err != nil {
			return nil, err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == 404 {
			resp.Body.Close()
			if skipOnNotFound {
				log.Println("optional dimension " + s.OrgName + "/" + dimensionKey + "/" + dimensionValue + " not found, skipping")
				return dimensionJsonMap, nil
			}
			log.Println("requested dimension not found response 404:" + resp.Request.URL.String())
			return nil, fmt.Errorf("dimension %s/%s/%s not found", s.OrgName, dimensionKey, dimensionValue)
		}

		if resp.StatusCode != 200 {
			resp.Body.Close()
			return nil, fmt.Errorf("request %s/%s/%s?workspace=%s failed with response: %v", s.OrgName, dimensionKey, dimensionValue, s.Workspace, resp.StatusCode)
		}
		defer resp.Body.Close()

		dimensionJsonBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("reading body response failed: %s", err)
		}

		var IacConsoleDBResponse IaCConsoleDBResponse
		err = json.Unmarshal(dimensionJsonBytes, &IacConsoleDBResponse)
		if err != nil {
			return nil, fmt.Errorf("error during unmarshal json response: %v", err)
		}

		if len(IacConsoleDBResponse.Dimensions) != 1 {
			return nil, fmt.Errorf("should be only one dimension in response")
		}
		if IacConsoleDBResponse.Error != "" {
			log.Println(IacConsoleDBResponse.Error)
		}
		dimensionJsonMap = IacConsoleDBResponse.Dimensions[0].DimData

	}

	return dimensionJsonMap, nil
}

func (s *State) ReportHistory(ctx context.Context, cmdToExec string, cmdArgs []string, cmdMainArg string, exitCode int) {
	if s.IacconsoleApiUrl == "" {
		// Only report to API if URL is configured
		return
	}

	var outputs map[string]interface{}

	if exitCode == 0 && (cmdMainArg == "apply" || cmdMainArg == "destroy") {
		// Run tofu output -json to gather outputs
		outputCmd := exec.Command(cmdToExec, "output", "-json")
		outputCmd.Dir = s.CmdWorkTempDir

		outputBytes, err := outputCmd.Output()
		if err == nil && len(outputBytes) > 0 {
			// Terraform outputs structure: {"name": {"sensitive": false, "type": "string", "value": "val"}}
			var tfOutputs map[string]interface{}
			if err := json.Unmarshal(outputBytes, &tfOutputs); err == nil {
				// We can either pass the entire tfOutputs or extract just the values
				// The API's HistoryRequestPost map[string]interface{} can handle the whole structure,
				// so let's pass it as is, or we can unpack values. Passing as is gives type and sensitive flags.
				outputs = tfOutputs
			} else {
				log.Printf("Failed to parse tf outputs: %v\n", err)
			}
		} else if err != nil {
			log.Printf("Failed to get tf outputs: %v\n", err)
		}
	}

	workspace := s.Workspace
	if workspace == "" {
		workspace = "master"
	}

	payload := map[string]interface{}{
		"cmdtoexec":  cmdToExec,
		"cmdargs":    cmdArgs,
		"cmdmainarg": cmdMainArg,
		"exitcode":   exitCode,
		"dimensions": s.ParsedDimensions,
		"outputs":    outputs,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal history payload: %v", err)
		return
	}

	url := fmt.Sprintf("%s/v1/history/%s/%s/%s", s.IacconsoleApiUrl, s.OrgName, workspace, s.UnitName)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(payloadBytes)))
	if err != nil {
		log.Printf("Failed to create request for history: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to report history to %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Failed to report history, status code %d: %s", resp.StatusCode, string(body))
	} else {
		log.Println("Successfully reported execution history to IaC Console API")
	}
}
