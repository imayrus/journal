package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/imayrus/journal/models"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		register()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

}

var id int = 0

func register() {
	file, err := os.Create("/tmp/journal/userList")
	checkNilErr(err)

	if _, err := os.Stat("/tmp/journal/userList"); err == nil {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("enter new 3username")
		u, _ := reader.ReadString('\n')
		newusername := strings.TrimSpace(u)

		logs := []models.LoginDetails{}
		f, err := os.Open("/tmp/journal/userList")
		checkNilErr(err)
		json.NewDecoder(f).Decode(&logs)

		for _, v := range logs {
			if v.UserName == newusername {
				fmt.Println("username already exist")
				break
			}

		}

		p := bufio.NewReader(os.Stdin)
		fmt.Println("enter password")
		pa, _ := p.ReadString('\n')
		newpassword := strings.TrimSpace(pa)

		id = id + 1

		log = append(log, models.LoginDetails{
			UserName: newusername,
			Password: newpassword,
			ID:       id,
		})

		fmt.Println("username and password added successfully \nlogin again ")

		usermarshal, err := json.Marshal(log)
		checkNilErr(err)

		_, err = file.Write(usermarshal)
		checkNilErr(err)

		LoggedInUser = newusername

		options()
	}
	// } else {
	// 	username, password := credentials()

	// 	log = (append(log, models.LoginDetails{
	// 		UserName: username,
	// 		Password: password,
	// 	}))

	// 	d, err := json.Marshal(log)
	// 	checkNilErr(err)
	// 	f.Write(d)

	// 	fmt.Println("User Registered")
	// 	LoggedInUser = username
	// 	options()
	// }
}

// data, err := ioutil.ReadFile("/tmp/journal/userList")
// checkNilErr(err)
// fmt.Println(string(data))

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
