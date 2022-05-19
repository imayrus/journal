package cmd

import (
	"bufio"
	"encoding/json"
	"strings"

	"fmt"
	"os"

	"github.com/imayrus/journal/models"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Existing User login",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

var LoggedInUser string
var log []models.LoginDetails

// func login() {

// reader := bufio.NewReader(os.Stdin)
// fmt.Println("enter username")
// u, _ := reader.ReadString('\n')
// username := strings.TrimSpace(u)

// 	for _, val := range log {
// 		if val.UserName == username {
// reader := bufio.NewReader(os.Stdin)
// fmt.Println("enter password")
// p, _ := reader.ReadString('\n')
// password := strings.TrimSpace(p)
// if val.Password == password {
// 	journal()
// } else {
// 	fmt.Println("wrong password")
// }

// 		} else {
// 			fmt.Println("Username doesn't exits.")
// 			break
// 		}

// 	}

// }
var UserId int

func login() {
	if _, err := os.Stat("/tmp/journal/userList"); err == nil {

		logs := []models.LoginDetails{}
		f, err := os.Open("/tmp/journal/userList")
		checkNilErr(err)
		defer f.Close()
		json.NewDecoder(f).Decode(&logs)

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("enter username")
		u, _ := reader.ReadString('\n')
		username := strings.TrimSpace(u)

		for _, v := range logs {
			if v.UserName == username {

				pa := bufio.NewReader(os.Stdin)
				fmt.Println("enter password")
				p, _ := pa.ReadString('\n')
				password := strings.TrimSpace(p)

				if v.Password == password {
					LoggedInUser = username
					fmt.Println("Welcome!", v.UserName)
					v.ID = UserId
					journal(UserId)
				} else {
					fmt.Println("Wrong Password")
					login()
				}
			} else {
				fmt.Println("Wrong Username ")
				register()
			}
		}

	} else {
		register()
	}
}

// func credentials() (string, string) {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println("enter username")
// 	u, _ := reader.ReadString('\n')
// 	username := strings.TrimSpace(u)

// 	pa := bufio.NewReader(os.Stdin)
// 	fmt.Println("enter password")
// 	p, _ := pa.ReadString('\n')
// 	password := strings.TrimSpace(p)

// 	return username, password

// }
