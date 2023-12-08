package rofi

import (
	"fmt"

	"github.com/bitfield/script"
)

func Dmenu(menu string) (string, error) {
	s, err := script.Echo(menu).Exec("rofi -dmenu").String()
	return s, err
}

func Message(msg string) error{
	fmt.Println(msg)
	_, err := script.Exec(fmt.Sprintf("rofi -e '%s'", msg)).String()
	return err
}
