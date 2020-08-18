package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	il "inlo/pkg/il"

	"github.com/spf13/cobra"
)

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
		var f *os.File = il.OpenOrCreateFile(currentTime)
		il.PrintTimeLine(f, currentTime, strings.Join(args, " "))
		defer f.Close()
	},
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
