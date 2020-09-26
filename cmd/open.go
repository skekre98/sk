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
	"os/exec"
	"strings"
	"errors"
	"path/filepath"
	"github.com/spf13/cobra"
)

// Function to compute Sørensen–Dice coefficient between two strings
func SorensenDice(stringOne, stringTwo string) float64 {
	i := strings.LastIndex(stringOne, "/")
	stringOne = stringOne[i+1:]


	firstBigrams := make(map[string]int)
	for i := 0; i < len(stringOne)-1; i++ {
		a := fmt.Sprintf("%c", stringOne[i])
		b := fmt.Sprintf("%c", stringOne[i+1])

		bigram := a + b

		var count int

		if value, ok := firstBigrams[bigram]; ok {
			count = value + 1
		} else {
			count = 1
		}

		firstBigrams[bigram] = count
	}

	var intersectionSize float64
	intersectionSize = 0
	for i := 0; i < len(stringTwo)-1; i++ {
		a := fmt.Sprintf("%c", stringTwo[i])
		b := fmt.Sprintf("%c", stringTwo[i+1])

		bigram := a + b

		var count int

		if value, ok := firstBigrams[bigram]; ok {
			count = value
		} else {
			count = 0
		}

		if count > 0 {
			firstBigrams[bigram] = count - 1
			intersectionSize = intersectionSize + 1
		}
	}

	return (2.0 * intersectionSize) / (float64(len(stringOne)) + float64(len(stringTwo)) - 2)
}

// Function to execute open command 
func openExecute(app string, destFile string) error {

	if app == "<>" && destFile == "<>" {
		return errors.New("missing application and file parameter")
	}

	app_map := make(map[string]string)
	app_map["sblm"] = "Sublime Text"
	app_map["vsc"] = "Visual Studio Code"
	app_map["adbl"] = "Adobe Lightroom"
	app_map["adbi"] = "Adobe Illustrator"
	app_map["adbp"] = "Adobe Photoshop CS6"

	files := []string{}
	err := filepath.Walk(".",
	    func(path string, info os.FileInfo, err error) error {
	    if err != nil {
	        return err
	    }
	    files = append(files, path)
	    return nil
	})
	if err != nil {
	    fmt.Println(err)
	}

	maxCoeff := 0.0
	path := ""
	for _, file := range files {
		currCoeff := SorensenDice(file, destFile)
		if currCoeff > maxCoeff {
			currCoeff = maxCoeff
			path = file
		}
	}

	app, exists := app_map[app]
	if exists == false {
		return errors.New("could not find desired app")
	}

	cmd := exec.Command("open", "-a", app, path)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "open a file or folder along the directory tree.",
	Long: `You'll never have to see a file to open it again |:>)(<:|`,
	Run: func(cmd *cobra.Command, args []string) {
		app, _ := cmd.Flags().GetString("app")
		file, _ := cmd.Flags().GetString("file")
		err := openExecute(app, file)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.
	openCmd.Flags().StringP("app", "a", "<>", "application for opening file")
	openCmd.Flags().StringP("file", "f", "<>", "file being opened")
}
