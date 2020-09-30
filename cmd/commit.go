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
	"os/exec"
	"errors"
	"github.com/spf13/cobra"
)

// Function to run git commands
func gitCommit(msg string, branch string) error {
	_, err := exec.Command("git", "add", ".").Output()
	if err != nil {
		return errors.New("ADD FAILED <!!!!!>")
	} else {
		fmt.Println("changes added <+>")
	}

	_, err = exec.Command("git", "commit", "-m", msg).Output()
	if err != nil {
		return errors.New("COMMIT FAILED <!!!!!>")
	} else {
		fmt.Println("changes commited <|>")
	}

	_, err = exec.Command("git", "push", "origin", branch).Output()
	if err != nil {
		return errors.New("PUSH FAILED <!!!!!>")
	} else {
		fmt.Println("changes pushed to branch %s <^>", branch)
	}

	return nil
}

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A command to push to Github",
	Long: `A replacement for git

Combines add, commit, and push into one command.`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, _ := cmd.Flags().GetString("msg")
		branch, _ := cmd.Flags().GetString("branch")
		err := gitCommit(msg, branch)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.
	commitCmd.Flags().StringP("msg", "m", ".", "commit message")
	commitCmd.Flags().StringP("branch", "b", "master", "repository branch")
}
