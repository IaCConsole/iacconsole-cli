package main

import "github.com/iacconsole/iacconsole-cli/cmd"

var (
	version string = "undefined"
)

func main() {
	cmd.SetVersionInfo(version)
	cmd.Execute()
}
