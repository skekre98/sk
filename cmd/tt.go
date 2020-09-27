/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// ttCmd represents the tt command
var ttCmd = &cobra.Command{
	Use:   "tt",
	Short: "Task Tracker",
	Long: `A tracker for analyzing and monitoring task completion.

The 'tt' command can be used to add or complete tasks. You can add
multiple tiers of difficulty to these tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		startTask, _ := cmd.Flags().GetString("start")
		endTask, _ := cmd.Flags().GetString("end")
		if startTask != "<>" {
			// TODO
			fmt.Println("task started")
		} else if endTask != "<>" {
			// TODO
			fmt.Println("task completed")
		} else if len(args) > 0 && args[0] == "ls" {
			// TODO
			fmt.Println("list called")
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(ttCmd)

	// Here you will define your flags and configuration settings.
	ttCmd.Flags().StringP("start", "s", "<>", "task being started")
	ttCmd.Flags().StringP("end", "e", "<>", "task being completed")
}
