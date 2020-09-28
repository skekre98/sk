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
	"os"
	"fmt"
	"time"
	"errors"
	"encoding/json"
	"io/ioutil"
	"github.com/spf13/cobra"
)

// task struct 
type Task struct {
	Text string `json:"Text"`
	Start time.Time `json:"Start"`
}

// Function to add task to list 
func addTask(task string) error {
	// Reading from file 
	home := os.Getenv("HOME")
	fileName := fmt.Sprintf("%s/.tt/tasks.json", home)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	
	// Preparing data 
	data := []Task{}
	json.Unmarshal(file, &data)
	for _, t := range data {
		if t.Text == task {
			return errors.New("Error: duplicate task already exists")
		}
	}
	newTask := &Task{
		Text: task,
		Start: time.Now(),
	}
	data = append(data, *newTask)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Writing to file 
	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Successfully added task")
	return nil
}

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
			err := addTask(startTask)
			if err != nil {
				fmt.Println(err.Error())
			}
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
