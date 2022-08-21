package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Initialize file for defman",
	Run: func(cmd *cobra.Command, args []string) {
		homedir := os.Getenv("HOME")
		filename := homedir + "/.config/defman/config.toml"
		resp, err := os.Stat(filename)
		if resp != nil {
			fmt.Println(filename + " is already exists")
		}
		if err != nil {
			fmt.Println(filename)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
