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
	"strings"
	// "errors"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	yt "github.com/knadh/go-get-youtube/youtube"
)

func youtubeDownload(url string) error {
	video, err := yt.Get(url)
	if err != nil {
		return err
	}

	options := &yt.Option{
		Rename: true,  // rename file using video title
		Resume: true,  // resume cancelled download
		Mp3:    true,  // extract audio to MP3
	}
	video.Download(0, "video.mp4", options)
	return nil
}

func youtubeSearch(query string) {
	query = strings.ReplaceAll(query, " ", "+")
	query = strings.Trim(query, " ")
	url := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", query)
	browser.OpenURL(url)
}

// youtubeCmd represents the youtube command
var youtubeCmd = &cobra.Command{
	Use:   "youtube",
	Short: "command to operate youtube from terminal",
	Long: `A command to interact with youtube such as searching,
downloading, and analyzing.`,
	Run: func(cmd *cobra.Command, args []string) {
		query, _ := cmd.Flags().GetString("search")
		link, _ := cmd.Flags().GetString("download")
		if query != "<>" {
			youtubeSearch(query)
		} else if link != "<>" {
			err := youtubeDownload(link)
			if err != nil {
				fmt.Println("Error:", err.Error())
			} 
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(youtubeCmd)

	// Here you will define your flags and configuration settings.
	youtubeCmd.Flags().StringP("search", "s", "<>", "search query")
	youtubeCmd.Flags().StringP("download", "d", "<>", "youtube download")
}