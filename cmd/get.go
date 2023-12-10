package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexpfx/go-pass_manager/pass"
	"github.com/alexpfx/go-pass_manager/rofi"
	"github.com/alexpfx/go-pass_manager/xdotool"
	"github.com/spf13/cobra"
)

var print bool
var debug bool

// menuCmd represents the get command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Aliases: []string{"m"},
	Short: "Mostra um menu com a lista de senhas disponíveis",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		userHome, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)

		}
		passStore := filepath.Join(userHome, ".password-store")

		list, _ := pass.List(passStore)

		s, err := rofi.Dmenu(strings.Join(list, "\n"))
		if err != nil {
			rofi.Message(fmt.Sprintf("Erro: %s", err))
			return
		}

		ps, err := pass.Show(s)
		if err != nil {
			rofi.Message(fmt.Sprintf("Erro: %s", err))
			return
		}
				
		_, err = xdotool.Type(ps, 55)
		if err != nil {
			rofi.Message(fmt.Sprintf("Erro: %s", err))
			return
		}

		if debug {
			log.Print(ps)
		}

		return

	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
	menuCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Imprime informações de debug")

}
