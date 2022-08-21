package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup config file for defman",
	Run: func(cmd *cobra.Command, args []string) {
		homedir := os.Getenv("HOME")
		dirname := homedir + "/.config/defman/"
		filename := "config.yaml"
		resp, _ := os.Stat(dirname + filename)
		if resp != nil {
			fmt.Println(dirname + filename + " is already exists👌")
			return
		}
		
		makeconfig(dirname, filename)
	},
}

func makeconfig(dirname, filename string) {
	if err := os.MkdirAll(dirname, 0755); err != nil {
		panic(err.Error())
	}

	_, err := os.Create(dirname + filename)
	if err != nil {
		panic(err.Error())
	}

	multi_line := []byte("---\n" + 
		"- Domain: NSGlobalDomain\n" + 
		"  Key: KeyRepeat\n" + 
		"  Type: int\n" + 
		"  Value: 1\n")
	ioutil.WriteFile(dirname + filename, multi_line, 0755)
	fmt.Println(dirname + filename + " was created😃")
	return
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
