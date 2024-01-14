/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/alexpfx/go-pass_manager/pass"
	"github.com/alexpfx/go-pass_manager/pm"
	"github.com/spf13/cobra"
)

var size int
var minUppercase int
var minLowercase int
var minDigits int
var minSpecials int
var letterCharset string
var numberCharset string
var specialCharset string

var insert bool
var force bool
var test bool

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Gera uma nova senha",
	Run: func(cmd *cobra.Command, args []string) {

		pm := pm.NewPass(letterCharset,
			numberCharset, specialCharset, minUppercase, minLowercase, minDigits, minSpecials, size)

		if insert {
			if len(args) < 1 {
				log.Fatal("Erro: a opção 'insert' deve ser acompanhada com o nome da senha. Use a sintaxe '-i <nome_da_senha>'.")
			}

			name := strings.TrimSpace(args[0])

			pass.Insert(pm.Generate(), name, force)

			return
		}
		fmt.Println(pm.Generate())

	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().IntVarP(&size, "size", "s", 12, "Tamanho da senha")
	genCmd.Flags().IntVarP(&minUppercase, "minUppercase", "C", 2, "Número mínimo de letras maíusculas")
	genCmd.Flags().IntVarP(&minLowercase, "minLowercase", "c", 2, "Número mínimo de letras minúsculas")
	genCmd.Flags().IntVarP(&minDigits, "minDigits", "d", 2, "Número mínimo de dígitos")
	genCmd.Flags().IntVarP(&minSpecials, "minSpecials", "x", 2, "Número mínimo caracteres especiais")

	genCmd.Flags().StringVar(&letterCharset, "letterCharset", "abcdefghijklmnopqrstuvxzwy", "Letras maísculas e minúsculas")

	genCmd.Flags().StringVar(&numberCharset, "numberCharset", "0123456789", "Números")

	genCmd.Flags().StringVar(&specialCharset, "specialCharset", "@#$:.!*-", "Caracteres especiais")
	genCmd.Flags().BoolVarP(&insert, "insert", "i", false, "Ao usar essa opção, uma nova senha será gerada e adicionada à sua coleção de senhas utilizando o comando 'pass'")
	genCmd.Flags().BoolVarP(&force, "force", "f", false, `Usando esta opção, você pode substituir uma senha existente pela nova senha gerada. 
				Se uma senha com o mesmo nome já existir, ela será automaticamente substituída pela nova senha.`)

	genCmd.Flags().BoolVarP(&test, "teste", "t", false, `flag para testes`)

}
