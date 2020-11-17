/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"bufio"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"strings"

	"github.com/navjordj/password_manager/crypto"
	"github.com/navjordj/password_manager/database"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/ssh/terminal"
)

// insertCmd represents the insert command
var insertPassword = &cobra.Command{
	Use:   "insert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using syour command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Website: ")
		website, _ := reader.ReadString('\n')

		fmt.Print("Enter your Password: ")
		userPassword, _ := terminal.ReadPassword(0)

		randomPassword := randomPassword(30)

		passwordEncrypted, _ := crypto.Encrypt(userPassword, []byte(randomPassword))

		res := database.Insert(string(passwordEncrypted), website, "test.db")
		if res == 1 {
			fmt.Println(fmt.Sprintf("\n %s is now inserted. Your password at %s is %s ", strings.TrimSuffix(website, "\n"), strings.TrimSuffix(website, "\n"), randomPassword))
		} else {
			fmt.Println(fmt.Sprintf("%s is already in the password manager", strings.TrimSuffix(website, "\n")))
		}

		//password_decrypted, _ := crypto.Decrypt([]byte("abc123"), password_encrypted)
		//fmt.Println(string(password_decrypted))

	},
}

func randomPassword(len int) string {
	res, err := password.Generate(len, 10, 10, true, true)
	if err != nil {
		log.Fatal(err)
	}

	return res

}

func init() {
	rootCmd.AddCommand(insertPassword)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
