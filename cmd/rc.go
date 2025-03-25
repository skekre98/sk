/*
Copyright © 2020 skekre98

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var remoteServer string

// getRemoteServer reads the remote server address from the properties file
func getRemoteServer() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %v", err)
	}

	configPath := filepath.Join(home, ".config", "sk_cli.properties")
	content, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("error reading properties file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "remote.server=") {
			return strings.TrimPrefix(line, "remote.server="), nil
		}
	}

	return "", fmt.Errorf("remote.server not found in properties file")
}

// setRemoteServer sets or updates the remote server address in the properties file
func setRemoteServer(server string) error {
	home, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %v", err)
	}

	configDir := filepath.Join(home, ".config")
	configPath := filepath.Join(configDir, "sk_cli.properties")

	// Create .config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %v", err)
	}

	// Read existing content if file exists
	var lines []string
	if content, err := os.ReadFile(configPath); err == nil {
		lines = strings.Split(string(content), "\n")
	}

	// Check if remote.server already exists
	serverSet := false
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "remote.server=") {
			lines[i] = fmt.Sprintf("remote.server=%s", server)
			serverSet = true
		}
	}

	// Add new remote.server if it doesn't exist
	if !serverSet {
		lines = append(lines, fmt.Sprintf("remote.server=%s", server))
	}

	// Write back to file
	content := strings.Join(lines, "\n")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing properties file: %v", err)
	}

	return nil
}

// rcCmd represents the rc command
var remoteCmd = &cobra.Command{
	Use:   "rc",
	Short: "Run command on remote server",
	Long:  `Run command on remote server based on configured host.`,
	Run: func(cmd *cobra.Command, args []string) {
		if remoteServer != "" {
			// Set new remote server
			if err := setRemoteServer(remoteServer); err != nil {
				fmt.Printf("Error setting remote server: %v\n", err)
				return
			}
			fmt.Printf("Remote server set to: %s\n", remoteServer)
			return
		}

		// Get current remote server
		server, err := getRemoteServer()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Remote server: %s\n", server)
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	// Add flags for setting the remote server
	remoteCmd.Flags().StringVarP(&remoteServer, "set-host", "s", "", "Set the remote server address")
}
