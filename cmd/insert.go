/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert new data",
	Long:  `This command inserts new data into the phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the data
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Not a valid name:", name)
			return
		}

		surname, _ := cmd.Flags().GetString("surname")
		if surname == "" {
			fmt.Println("Not a valid surname:", surname)
			return
		}

		tel, _ := cmd.Flags().GetString("telephone")
		if tel == "" {
			fmt.Println("Not a valid telephone:", tel)
			return
		}

		t := strings.ReplaceAll(tel, "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		err := insert(name, surname, t)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("name", "n", "", "name value")
	insertCmd.Flags().StringP("surname", "s", "", "surname value")
	insertCmd.Flags().StringP("telephone", "t", "", "telephone value")
}

func insert(name, surname, tel string) error {
	// If it already exists, do not add it
	_, ok := index[tel]
	if ok {
		return fmt.Errorf("%s already exists", tel)
	}

	data = append(data, Entry{
		Name:       name,
		Surname:    surname,
		Tel:        tel,
		LastAccess: strconv.FormatInt(time.Now().Unix(), 10),
	})

	// Update the index
	createIndex()

	return saveJSONFile(filepath)
}
