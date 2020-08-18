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
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	layoutISO  = "2006-01-02"
	lineFormat = "15:04:05"
)

type incidentLog *os.File

// tlCmd represents the tl command
var tlCmd = &cobra.Command{
	Use:   "tl",
	Short: "Append a timeline item",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tl called")

		location, err := time.LoadLocation("America/Toronto")
		if err != nil {
			log.Fatal(err)
		}

		currentTime := time.Now().In(location)
		var f *os.File = openOrCreateFile(currentTime)
		printTimeLine(f, currentTime, strings.Join(args, " "))
		defer f.Close()
	},
}

func printTimeLine(file *os.File, currentTime time.Time, message string) {
	line := fmt.Sprintf("%s|TIMELINE - %s\n", currentTime.Format(lineFormat), message)
	if _, err := file.WriteString(line); err != nil {
		log.Fatal(err)
	}
}

func openOrCreateFile(date time.Time) *os.File {

	fileName := fmt.Sprintf("%s.md", date.Format(layoutISO))
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	var osPermissions os.FileMode = 0664

	f, err := os.OpenFile(fileName, flags, osPermissions)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func init() {
	rootCmd.AddCommand(tlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
