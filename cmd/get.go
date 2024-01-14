package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexpfx/go-pass_manager/pass"
	"github.com/alexpfx/go-pass_manager/pm"
	"github.com/alexpfx/go-pass_manager/rofi"
	"github.com/alexpfx/go-pass_manager/wtype"
	"github.com/spf13/cobra"
)

var print bool
var debug bool

// menuCmd represents the get command
var menuCmd = &cobra.Command{
	Use:     "menu",
	Aliases: []string{"m"},
	Short:   "Mostra um menu com a lista de senhas disponíveis",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {

		dmenuTool := getDmenuTool()
		userHome, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)

		}
		passStore := filepath.Join(userHome, ".password-store")

		list, _ := pass.List(passStore)

		s, err := dmenuTool.Dmenu(strings.Join(list, "\n"))
		if err != nil {
			dmenuTool.Message(fmt.Sprintf("Erro ao mostrar menu: %s", err.Error()))
			return
		}

		ps, err := pass.Show(s)
		if err != nil {
			dmenuTool.Message(fmt.Sprintf("Erro ao obter senha: %s", err.Error()))
			return
		}

		ttool := getTypeTool()

		_, err = ttool.Type(ps, 100)
		if err != nil {
			dmenuTool.Message(fmt.Sprintf("Erro ao digitar senha: %s", err.Error()))
			return
		}

		if debug {
			log.Print(ps)
		}

		return

	},
}

func getTypeTool() pm.Typist {
	// return xdotool.New()
	return &wtype.WType{}
}

func getDmenuTool() pm.Menu {
	return &rofi.Rofi{}
}

func init() {
	rootCmd.AddCommand(menuCmd)
	menuCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Imprime informações de debug")

}
