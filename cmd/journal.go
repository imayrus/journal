package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/imayrus/journal/models"
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

var entry string

func journal() {

	var scanner = bufio.NewScanner((os.Stdin))

	for {
		fmt.Println(`Choose Option
		1. Add New Journal Entry
		2. View Journal Entry
		3. Update Journal
		4. Delete Journal
		5. Logout
		6. Exit`)
		scanner.Scan()
		choice, _ := strconv.Atoi((scanner.Text()))

		switch choice {
		case 1:
			AddNewEntry()
		case 2:
			ViewEntry()
		case 3:
			UpdateEntry()
		case 4:
			DeleteEntry()
		case 5:
			Logout()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice")
		}
	}
}

func AddNewEntry() {
	if isLoggedIn() {
		if _, err := os.Stat("/tmp/journal/entries/" + LoggedInUser + "/journal"); err == nil {
			entry := getNewEntry()
			fmt.Println(entry)
		} else {
			f, err := os.Create("/tmp/journal/entries/" + LoggedInUser + "/journal")
			checkNilErr(err)
			entry := getNewEntry()
			e := []byte(entry)
			f.Write(e)
		}
	} else {
		fmt.Println("login first")
		Execute()
	}

}

func getNewEntry() string {
	var entries []string
	entry = ""
	fmt.Println("\nAdd your new entry. You can write in multiple lines. Press Ctrl+Enter on a new line when done. \n")
	scn := bufio.NewScanner(os.Stdin)
	for {
		for scn.Scan() {
			line := scn.Text()
			if len(line) == 1 {
				if line[0] == '\x1D' {
					break
				}
			}
			entries = append(entries, line)
		}
		if err := scn.Err(); err != nil {
			fmt.Println(os.Stderr, err)
			break
		}
		if len(entries) == 0 {
			break
		}
	}

	for _, value := range entries {
		entry = entry + value + "\n"
	}

	return entry
}

func ViewEntry() {
	if isLoggedIn() {
		if _, err := os.Stat("/tmp/journal/entries/" + LoggedInUser); err == nil {
			entry := []models.Entry{}
			f, err := os.Open("/tmp/journal/entries/" + LoggedInUser)
			checkNilErr(err)
			defer f.Close()
			json.NewDecoder(f).Decode(&entry)

			fmt.Println(entry)
		}

	}
	os.Exit(0)
}

func UpdateEntry() {
	if isLoggedIn() {
		if _, err := os.Stat("/tmp/journal/entries" + LoggedInUser + "/journal"); err == nil {
			entry := []models.Entry{}
			f, err := os.Open("/tmp/journal/entries" + LoggedInUser + "/journal")
			checkNilErr(err)
			defer f.Close()
			json.NewDecoder(f).Decode(&entry)

			var entries []string
			scn := bufio.NewScanner(os.Stdin)

			for scn.Scan() {
				entry := scn.Text()
				if len(entry) == 1 {
					if entry[0] == '\x1D' {
						break
					}
				}
				entries = append(entries, entry)

				if err := scn.Err(); err != nil {
					fmt.Println(os.Stderr, err)
					break
				}
				if len(entries) == 0 {
					break
				}
			}
		}

	}

}

func DeleteEntry() {
	if isLoggedIn() {
		if err := os.Truncate("/tmp/journal/entries"+LoggedInUser+"/journal", 0); err != nil {
			fmt.Printf("Failed to truncate: %v", err)
		}
	}
	os.Exit(0)
}

func Logout() {
	fmt.Println("logout successfully")
	options()
}

func isLoggedIn() bool {
	if LoggedInUser != "" {
		return true
	}
	return false
}
