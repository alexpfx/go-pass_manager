package wofi

import (
	"fmt"
	"log"

	"github.com/bitfield/script"
)

func (w Wofi) Dmenu(menu string) (string, error) {
	log.Println(menu)
	s, err := script.Echo(menu).Exec("wofi -dmenu").String()
	return s, err
}

func (w Wofi) Message(msg string) error {
	fmt.Println(msg)
	_, err := script.Exec(fmt.Sprintf("wofi -e '%s'", msg)).String()
	return err
}

type Wofi struct{}
