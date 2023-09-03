package main

import (
	"fmt"
	"os"
)

var (
	appName = "httpd"
)

func configFile() (string, error) {
	path := os.Getenv("HTTPD_CONFIG_FILE")
	if path != "" {
		return path, nil
	}

	userCfg, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("can't get user config dir - %w", err)
	}

	path = fmt.Sprintf("%s/%s/config.toml", userCfg, appName)
	return path, nil
}

func main() {
	fmt.Println(configFile())
}
