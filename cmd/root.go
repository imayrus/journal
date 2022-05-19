package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "Basic Journal CLI App",
	Run: func(cmd *cobra.Command, args []string) {
		options()
	},
}

func init() {
	rootCmd.AddCommand(rootCmd)

}

func options() {
	fmt.Println("Welcome to the journal. \nChoose the below option: \n\n1. Login \n2. Register \n3. Exit")
	option := bufio.NewReader(os.Stdin)
	choice, _ := option.ReadString('\n')

	choice = strings.Trim(choice, "\n")
	if choice == "1" {
		login()
	} else if choice == "2" {
		register()
	} else if choice == "3" {
		os.Exit(0)
	} else {
		fmt.Println("Invalid choice")
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
