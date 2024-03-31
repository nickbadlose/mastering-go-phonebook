/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for the number",
	Long: `search whether a telephone number exists in the
	phone book application or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		if key == "" {
			fmt.Println("Not a valid key:", key)
			return
		}

		t := strings.ReplaceAll(key, "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
		}

		result := search(t)
		if result == nil {
			fmt.Println("Entry not found:", t)
			return
		}
		fmt.Println(*result)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().String("key", "", "Key to search")
}

func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}

	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &data[i]
}
