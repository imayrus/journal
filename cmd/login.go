package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

type LoginDetails struct {
	UserName string
	Password string
}

var log []LoginDetails

func login() {
	log = append(log, LoginDetails{"abc", "def"})
	log = append(log, LoginDetails{"surya", "prakash"})

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter username")
	u, _ := reader.ReadString('\n')
	username := strings.TrimSpace(u)

	for _, val := range log {
		if val.UserName == username {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("enter password")
			p, _ := reader.ReadString('\n')
			password := strings.TrimSpace(p)
			if val.Password == password {
				journal()
			} else {
				fmt.Println("Enter wrong password")
			}

		}
	}
}

func register() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter username")
	username, _ := reader.ReadString('\n')
	newusername := strings.TrimSpace(username)
	password := bufio.NewReader(os.Stdin)
	fmt.Println("enter password")
	pa, _ := password.ReadString('\n')
	newpassword := strings.TrimSpace(pa)

	log = append(log, LoginDetails{newusername, newpassword})
	fmt.Println("username and password added successfully \nlogin again ")
	login()
}
