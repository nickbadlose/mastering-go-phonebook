/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phonebook",
	Short: "A phone book application",
	Long:  `This is a Phone Book application that uses JSON records.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(setJSONFile())

	err := readJSONFile(filepath)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(1)
	}

	createIndex()

	cobra.CheckErr(rootCmd.Execute())
}

func init() {}

const (
	cfgPhonebook = "PHONEBOOK"
)

var (
	data     = phonebook{}
	index    = map[string]int{}
	filepath = "/Users/nick/projects/mastering-go-phonebook/data.json"
)

type Entry struct{ Name, Surname, Tel, LastAccess string }

type phonebook []Entry

func (p phonebook) Len() int { return len(p) }
func (p phonebook) Less(i, j int) bool {
	if p[i].Surname == p[j].Surname {
		return p[i].Name < p[j].Name
	}
	return p[i].Surname < p[j].Surname
}
func (p phonebook) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func Serialize(slice interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(slice)
}

func DeSerialize(slice interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(slice)
}

func setJSONFile() error {
	fp := os.Getenv(cfgPhonebook)
	if fp != "" {
		filepath = fp
	}

	_, err := os.Stat(filepath)
	// If error is not nil, it means that the file does not exist
	if err != nil {
		fmt.Println("Creating", filepath)
		f, err := os.Create(filepath)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return err
		}
		f.Close()
	}

	fileInfo, err := os.Stat(filepath)
	// Is it a regular file?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(filepath, "not a regular file!")
		return fmt.Errorf("%s is not a regular file", filepath)
	}

	return nil
}

func saveJSONFile(fp string) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	return Serialize(&data, f)
}

func readJSONFile(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	return DeSerialize(&data, f)
}

func matchTel(t string) bool {
	rgx := regexp.MustCompile(`^\d+$`)
	return rgx.Match([]byte(t))
}

func createIndex() {
	index = make(map[string]int, len(data))
	for i, entry := range data {
		index[entry.Tel] = i
	}
}
