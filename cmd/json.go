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
	"bytes"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/cobra"
)

// Function to indent json data
func jsonBytes(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
}


// Function to print json file 
func printJson(fileName string) error {
	jsonFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	b := []byte(jsonFile)
	b, err = jsonBytes(b)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", b)
	return nil
}

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "<~>",
	Long: `command to print a json file in indented format

file must be in current directory tree`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		if file == "<>" {
			fmt.Println("Error: missing file parameter")
		} else {
			err := printJson(file)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	// Here you will define your flags and configuration settings.
	jsonCmd.Flags().StringP("file", "f", "<>", "json file")
}
