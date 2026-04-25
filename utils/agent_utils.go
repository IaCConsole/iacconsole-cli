package utils

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// AgentMessage is the base structure for all WebSocket messages
type AgentMessage struct {
	Type string `json:"type"`
}

// AgentRegister sent by agent upon connection
type AgentRegister struct {
	AgentMessage
	AgentID string `json:"agentId"`
	Version string `json:"version"`
	OS      string `json:"os"`
	Arch    string `json:"arch"`
}

// AgentCommand received by agent from server
type AgentCommand struct {
	AgentMessage
	ID         string          `json:"id"`
	Action     string          `json:"action"` // init, plan, apply, destroy
	Org        string          `json:"org"`
	Unit       string          `json:"unit"`
	Dimensions []DimensionPair `json:"dimensions"`
	Workspace  string          `json:"workspace,omitempty"`
	ExtraArgs  []string        `json:"extraArgs,omitempty"`
}

type DimensionPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AgentOutput sent by agent to server
type AgentOutput struct {
	AgentMessage
	CommandID string `json:"commandId"`
	Stream    string `json:"stream"` // stdout, stderr
	Data      string `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

// AgentComplete sent by agent when command finishes
type AgentComplete struct {
	AgentMessage
	CommandID string `json:"commandId"`
	ExitCode  int    `json:"exitCode"`
	Error     string `json:"error,omitempty"`
}

// AgentPing received from browser via server
type AgentPing struct {
	AgentMessage
}

// AgentPong sent back to browser via server
type AgentPong struct {
	AgentMessage
	Timestamp int64 `json:"timestamp"`
}

// AgentError sent by server to agent when there's an error
type AgentError struct {
	AgentMessage
	Error string `json:"error"`
}

// ParseAPIURL parses the IACCONSOLE_API_URL environment variable and handles IACCONSOLE_TOKEN
func GetAPIURL() (wsURL string, apiURLHost string, authHeader string, err error) {
	apiUrl := os.Getenv("IACCONSOLE_API_URL")
	token := os.Getenv("IACCONSOLE_TOKEN")

	if token == "" {
		return "", "", "", fmt.Errorf("IACCONSOLE_TOKEN is required. Please set it in your environment")
	}

	// Use default API URL if empty
	if apiUrl == "" {
		apiUrl = "https://api.iacconsole.com"
	}

	// Remove trailing slash
	apiUrl = strings.TrimRight(apiUrl, "/")

	if !strings.HasPrefix(apiUrl, "http") {
		return "", "", "", fmt.Errorf("IACCONSOLE_API_URL must start with http:// or https://")
	}

	u, err := url.Parse(apiUrl)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to parse IACCONSOLE_API_URL: %v", err)
	}

	scheme := "wss"
	if u.Scheme == "http" || strings.HasPrefix(u.Host, "localhost:") {
		scheme = "ws"
	}

	wsURL = fmt.Sprintf("%s://%s/v1/ws/agent", scheme, u.Host)

	// Token Authentication is mandatory
	authHeader = "Bearer " + token
	return wsURL, apiUrl, authHeader, nil
}
