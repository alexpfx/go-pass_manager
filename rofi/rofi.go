package rofi

import (
	"fmt"
	"strings"

	"github.com/alexpfx/linux_wrappers/linux"
)

func Dmenu(menu string) (string, error) {
	r := linux.NewDMenu("Escolha a senha")
	s, err := r.Run(menu)
	return strings.TrimSpace(s), err
}

func Message(msg string) (string, error) {
	fmt.Println(msg)
	r := linux.NewMessageMenu(msg)
	return r.Run(msg)
}
