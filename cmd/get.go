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
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/navjordj/password_manager/crypto"
	"github.com/navjordj/password_manager/database"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"

	"bufio"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		db, _ := sql.Open("sqlite3", "test.db")

		fmt.Println("get called")

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Website: ")
		website, _ := reader.ReadString('\n')

		fmt.Print("Enter your Password: ")
		userPassword, _ := terminal.ReadPassword(0)

		if database.CheckInDatabase(website, db) == true {
			res := database.Get(website, db)
			decrypted, err := crypto.Decrypt([]byte(userPassword), []byte(res))

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(decrypted))
		} else {

		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
