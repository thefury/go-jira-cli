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

// commentListCmd represents the commentList command
var commentListCmd = &cobra.Command{
	Use:   "list",
	Short: "list comments",
	Long:  `list comments for a given issue`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString("issue")
		if err != nil {
			panic(err)
		}

		issue, err := GetIssue(key)
		if err != nil {
			panic(err)
		}

		formatter := JSONFormatter{Data: issue.Fields.Comments}
		jsonData := formatter.Format()

		fmt.Println(jsonData)
	},
}

func init() {
	commentCmd.AddCommand(commentListCmd)
	commentListCmd.Flags().StringP("issue", "i", "", "jira key or issue id (required)")
	commentListCmd.MarkFlagRequired("issue")
}
