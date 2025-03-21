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
	"net/url"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

var port int

// lhCmd represents the lh command
var lhCmd = &cobra.Command{
	Use:   "lh [port]",
	Short: "Open localhost in your default browser",
	Long: `Open localhost in your default browser.
If no port is provided, defaults to port 8080.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			p, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Error: Invalid port number '%s'\n", args[0])
				return
			}
			port = p
		}

		// Create the localhost URL
		url := url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("localhost:%d", port),
		}

		// Open the URL in the default browser based on the OS
		var err error
		switch runtime.GOOS {
		case "darwin":
			err = exec.Command("open", url.String()).Start()
		case "windows":
			err = exec.Command("cmd", "/c", "start", url.String()).Start()
		default: // linux and others
			err = exec.Command("xdg-open", url.String()).Start()
		}

		if err != nil {
			fmt.Printf("Error opening browser: %v\n", err)
			return
		}

		fmt.Printf("Opening http://localhost:%d in your browser...\n", port)
	},
}

func init() {
	rootCmd.AddCommand(lhCmd)

	// Set default port to 8080
	port = 8080
}
