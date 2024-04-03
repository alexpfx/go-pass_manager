package cmd

import (
	"fmt"
	"github.com/alexpfx/linux_wrappers/linux"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexpfx/go-pass_manager/pass"
	"github.com/alexpfx/go-pass_manager/rofi"
	"github.com/spf13/cobra"
)

var list bool = false
var debug bool

// menuCmd represents the get command
var menuCmd = &cobra.Command{
	Use:     "menu",
	Aliases: []string{"m"},
	Short:   "Mostra um menu com a lista de senhas disponíveis",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {
		userHome, err := os.UserHomeDir()
		fmt.Println(userHome)
		if err != nil {
			log.Fatal(err)

		}
		passStore := filepath.Join(userHome, ".password-store")

		passlist, err := pass.List(passStore)
		if err != nil {
			log.Fatal(err)
		}

		jlist := strings.Join(passlist, "\n")
		if list {
			fmt.Println("jlist")
			fmt.Println(jlist)
			return
		}
		ps, err := rofi.Dmenu(jlist)
		if err != nil {
			rofi.Message(fmt.Sprintf("Erro ao mostrar menu: %s", err.Error()))
			return
		}

		out, err := pass.Show(ps)
		if err != nil {
			rofi.Message(fmt.Sprintf("Erro ao obter senha: %s", err.Error()))
			return
		}
		ttool := linux.NewWType(linux.WTypeBuilder{
			DelayBetweenKeyStrokes: "50",
			DelayBeforeKeyStrokes:  "200",
		})

		_, err = ttool.Run(out)

		if err != nil {
			rofi.Message(fmt.Sprintf("Erro ao digitar senha: %s", err.Error()))
			return
		}

		if debug {
			log.Print(out)
		}

		return

	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
	menuCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Imprime informações de debug")
	menuCmd.Flags().BoolVar(&list, "list", false, "Apenas lista o menu na stdout")

}
