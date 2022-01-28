package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var journalCmd = &cobra.Command{
	Use:   "journal",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(journalCmd)

}

func journal() {
	var scanner = bufio.NewScanner((os.Stdin))

	for {
		fmt.Println(`Choose Option
		1. Add Journal
		2. View Journal
		3. Exit`)
		scanner.Scan()
		choice, _ := strconv.Atoi((scanner.Text()))

		switch choice {
		case 1:
			addentry()
		case 2:
			viewentry()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice")
		}
	}
}

func addentry() {
	fmt.Println("add")
	os.Exit(0)
}

func viewentry() {
	fmt.Println("view")
	os.Exit(0)
}
