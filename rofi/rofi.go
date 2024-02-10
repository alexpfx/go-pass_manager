package rofi

import (
	"fmt"

	"github.com/alexpfx/linux_wrappers/linux"
)

func Dmenu(menu string) (string, error) {
	r := linux.NewDMenu("Escolha a senha")
	return r.Run(menu)
}

func Message(msg string) (string, error) {
	fmt.Println(msg)
	r := linux.NewMessageMenu(msg)
	return r.Run(msg)
}
