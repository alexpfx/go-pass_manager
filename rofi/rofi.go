package rofi

import (
	"fmt"
	"github.com/alexpfx/linux_wrappers/wrappers/rofi"
	"strings"
)

func Dmenu(menu string) (string, error) {
	r := rofi.New("Escolha a Senha")
	out, err := r.ShowDMenu(menu)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), err
}

func Message(msg string) (string, error) {
	fmt.Println(msg)
	r := rofi.NewMessage(msg)
	return r.ShowDMenu(msg)
}
