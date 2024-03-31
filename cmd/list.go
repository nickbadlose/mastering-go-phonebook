/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all entries",
	Long:  `This command lists all entries in the phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func list() {
	sort.Sort(data)
	pretty, err := PrettifyJSONStream(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pretty)
	fmt.Printf("%d records in total.\n", len(data))
}

func PrettifyJSONStream(d interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(d)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
