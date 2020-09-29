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
	Level int `json:"Level"`
}

// finished Task struct 
type FinishedTask struct {
	Text string `json:"Text"`
	Duration time.Duration `json:"Duration"`
	Level int `json:"Level"`	
}

// Function to mark task as completed 
func markComplete(finished Task) error {
	// Reading from file 
	home := os.Getenv("HOME")
	fileName := fmt.Sprintf("%s/.tt/completed.json", home)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	
	// Preparing data 
	data := []FinishedTask{}
	json.Unmarshal(file, &data)
	newFinish := &FinishedTask{
		Text: finished.Text,
		Duration: time.Now().Sub(finished.Start),
		Level: finished.Level,
	}
	fmt.Println(newFinish.Duration)
	data = append(data, *newFinish)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Writing to file 
	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Function to remove task from task file 
func popTask(filename string, task string) ([]Task, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	// Preparing data 
	taskArray := []Task{}
	json.Unmarshal(file, &taskArray)	

	pos := -1
	for i, t := range taskArray {
		if t.Text == task {
			pos = i
            break
		}
	}
	if pos == -1 {
		return nil, errors.New("task does not exist")
	}

	// Mark task as completed 
	completedTask := taskArray[pos]
	err = markComplete(completedTask)
	if err != nil {
		return nil, err
	}

	// Return updated list
	copy(taskArray[pos:], taskArray[pos+1:]) 
	taskArray[len(taskArray)-1] = Task{}
	taskArray = taskArray[:len(taskArray)-1]

	return taskArray, nil
}

// Function to migrate task to completed list 
func migrateTask(task string) error {
	// Reading from file 
	home := os.Getenv("HOME")
	taskFile := fmt.Sprintf("%s/.tt/tasks.json", home)
	taskArray, err := popTask(taskFile, task)
	if err != nil {
        return err
    }

	dataBytes, err := json.Marshal(taskArray)
	if err != nil {
		return err
	}

	// Writing to file 
	err = ioutil.WriteFile(taskFile, dataBytes, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Successfully migrated task")
	return nil
}

// Function to add task to list 
func addTask(task string, level int) error {
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
		Level: level,
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
		endTask, _ := cmd.Flags().GetString("complete")
		taskLevel, _ := cmd.Flags().GetInt("level")
		if startTask != "<>" {
			err := addTask(startTask, taskLevel)
			if err != nil {
				fmt.Println("Error:", err.Error())
			}
		} else if endTask != "<>" {
			err := migrateTask(endTask)
			if err != nil {
                fmt.Println("Error:", err.Error())
            }
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
	ttCmd.Flags().IntP("level", "l", 1, "difficulty level")
	ttCmd.Flags().StringP("complete", "c", "<>", "task being completed")
}
