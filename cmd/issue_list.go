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

// issueListCmd represents the issueList command
var issueListCmd = &cobra.Command{
	Use:   "list",
	Short: "list jira issues",
	Long:  `list Jira issues for a given project`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString("project")
		if err != nil {
			panic(err)
		}

		issues, err := GetIssueList(key)
		if err != nil {
			panic(err)
		}

		formatter := JSONFormatter{Data: issues}
		jsonData := formatter.Format()

		fmt.Println(jsonData)
	},
}

func init() {
	issueCmd.AddCommand(issueListCmd)

	issueListCmd.Flags().StringP("project", "p", "", "project key (required)")
	issueListCmd.MarkFlagRequired("project")
}
