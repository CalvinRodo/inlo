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
	"inlo/consts"
	"inlo/pkg/il"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// attachCmd represents the attach command
var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fileName := args[0]

		from, err := os.Open(filepath.Join(fileName))
		if err != nil {
			panic(err)
		}
		defer from.Close()

		toPath := filepath.Join(viper.GetString(consts.LOGDIR), fileName)
		flags := os.O_RDWR | os.O_CREATE
		to, err := os.OpenFile(toPath, flags, 0666)
		if err != nil {
			panic(err)
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		if err != nil {
			panic(err)
		}

		il.PrintLine("FILEATTACHED", fmt.Sprintf("[%s](%s)", fileName, toPath))
		fmt.Printf("File %s attached\n", fileName)
	},
}

func init() {
	rootCmd.AddCommand(attachCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// attachCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// attachCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
