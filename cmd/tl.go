package cmd

import (
	"fmt"
	"inlo/pkg/il"
	"strings"

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
		il.PrintLine("TIMELINE", strings.Join(args, " "))
		fmt.Println("Event added")
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
