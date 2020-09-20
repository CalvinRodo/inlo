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
	"bufio"
	"inlo/halt"
	"inlo/pkg/attach"
	"inlo/pkg/il"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read the contents of a pipe into a file",
	Long: `Read the contents of a pipe into the file passed as the first argument. 
ex: 

inlo read foo.txt << EOF 
Foo
Bar
EOF

and 

ps -ax | inlo read processes.tx

`,
	Run: func(cmd *cobra.Command, args []string) {

		//fileName := args[0]

		info, err := os.Stdin.Stat()
		halt.IfErr(err)

		log.Printf("PipeMode: %s", info.Mode().String())

		if info.Mode()&os.ModeNamedPipe == 0 {
			panic("the command is intendended to work with pipes.")
		}

		reader := bufio.NewReader(os.Stdin)

		fileName := args[0]
		attach.CopyFile(fileName, reader)
		il.PrintLine("STDIN", attach.MdFileLink(fileName))

	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
