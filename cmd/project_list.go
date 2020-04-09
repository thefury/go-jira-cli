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

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "list projects",
	Long: `List jira projects that you have access to. This list can
be filtered.`,

	Run: func(cmd *cobra.Command, args []string) {
		projects, err := GetProjectList()
		if err != nil {
			panic(err)
		}

		formatter := JSONFormatter{Data: projects}
		jsonData := formatter.Format()

		fmt.Println(jsonData)

	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
