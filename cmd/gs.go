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
	"context"
	"strconv"
	"github.com/spf13/cobra"
	google "github.com/rocketlaunchr/google-search"
	"github.com/jedib0t/go-pretty/table"
)

// Function to run search query
func runQuery(query string) error {
	ctx := context.Background()
	var results []google.Result
	results, err := google.Search(ctx, query)
	if err != nil {
		return err
	}
	
	// Creating table of results 
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"", "Title", "URL"})
	for i := 0; i < 5; i++ {
		r := results[i]
		t.AppendRow([]interface{}{r.Rank, r.Title, r.URL})
	}
	t.Render()

	fmt.Println("Which result would you like to route to?")
	fmt.Print("	(1-5)/n: ")
	var resp string
	fmt.Scanln(&resp)
	if resp == "n" {
		fmt.Println("Routing cancelled")
		return nil
	}

	index, err := strconv.Atoi(resp)
	if err != nil {
		return err
	}
	fmt.Println(results[index-1].URL)
	return nil
}

// gsCmd represents the gs command
var gsCmd = &cobra.Command{
	Use:   "gs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		query, _ := cmd.Flags().GetString("query")
		if query == "<>" {
			fmt.Println("Query was left empty")
		} else {
			err := runQuery(query)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(gsCmd)

	// Here you will define your flags and configuration settings.
	gsCmd.Flags().StringP("query", "q", "<>", "search query")
}