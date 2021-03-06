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
	"github.com/charmbracelet/glamour"
	"inlo/cmd/attach"
	"inlo/cmd/folder"
	"inlo/cmd/halt"
	"inlo/cmd/il"
	"inlo/cmd/settings"
	"strings"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("view called")

		logFolder := folder.ExpandDir(settings.Settings.LogPath)
		logFileName := il.FileNameForToday(logFolder)
		contents := attach.ReadFile(logFileName)

		log := renderFile(contents)
		out, err := glamour.Render(log, "dark")
		halt.IfErr(err)
		fmt.Print(out)
	},
}

func mapStrings(strings []string, f func(string) string) []string {

	retval := make([]string,0)

	for _, s := range strings {
		retval = append(retval, f(s))
	}
	return retval
}


func renderFile(c []string) string {
	toList := func(s string) string {
		return "- " + s
	}

	return strings.Join(mapStrings(c, toList) , "\n")
}



func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
