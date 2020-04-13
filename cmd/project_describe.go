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

// projectShowCmd represents the projectShow command
var projectDescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a project",
	Long: `Describe a single project. 

Requires a project to be given.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString("project")
		if err != nil {
			panic(err)
		}

		project, err := GetProject(key)
		if err != nil {
			panic(err)
		}

		// TODO most of this can be pulled ito a class
		formatter := JSONFormatter{Data: project}
		jsonData := formatter.Format()

		fmt.Println(jsonData)
	},
}

func init() {
	projectCmd.AddCommand(projectDescribeCmd)

	projectDescribeCmd.Flags().StringP("project", "p", "", "project key (required)")
	projectDescribeCmd.MarkFlagRequired("project")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectShowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectShowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
