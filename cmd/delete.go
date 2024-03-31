/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an entry",
	Long:  `delete an entry from the phone book application.`,
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

		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().String("key", "", "Key to delete")
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}

	data = append(data[:i], data[i+1:]...)

	// Update the index - key does not exist anymore
	// This is pointless since the index is created at the start of each run,
	// also you would need to update the keys of all the entries after the deleted one in the slice. But hey nvm for now.
	delete(index, key)

	return saveJSONFile(filepath)
}
