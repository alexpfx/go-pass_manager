/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/alexpfx/go-pass_manager/backup"
	"github.com/alexpfx/go-pass_manager/pass"
	"github.com/spf13/cobra"
)

var export bool
var load bool

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Armazena e recupera backups",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if export {
			filename := createTempFile()
			err := backup.Export(filename)
			if err != nil {
				log.Fatal(err)
			}

		}

		if load {
			if len(args) > 0 {
				f, err := os.Open(args[0])
				if err != nil {
					log.Fatal(err)
				}
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					splitted := strings.Split(scanner.Text(), ":")
					pn := splitted[0]
					p := splitted[1]
					err := pass.Insert(p, pn, false)
					if err != nil {
						log.Print(err)
						continue
					}
				}

			}

		}

	},
}

func createTempFile() string {
	// Create a temporary file
	f, err := os.CreateTemp(".", ".senhas_salvas-")

	if err != nil {
		log.Fatal(err)
	}

	f.Close()
	return f.Name()
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	backupCmd.Flags().BoolVar(&export, "export", false, "Faz backup de todas as senhas")
	backupCmd.Flags().BoolVar(&load, "load", false, "Carrega todas as senhas do backup")

}
