package cmd

import (
	"fmt"
	"inlo/pkg/il"
	"strings"

	"github.com/spf13/cobra"
)

// icCmd represents the ic command
var icCmd = &cobra.Command{
	Use:   "ic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ic := strings.Join(args, " ")
		il.PrintLine("INCIDENT COMMANDER", ic)
		fmt.Printf("Incident Command %s Added\n", ic)
	},
}

func init() {
	rootCmd.AddCommand(icCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// icCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// icCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
