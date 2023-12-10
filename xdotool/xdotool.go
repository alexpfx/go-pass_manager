package xdotool

import (
	"fmt"

	"github.com/bitfield/script"
)

func Type(text string, delayMs int) (string, error){
	return script.Exec(fmt.Sprintf("xdotool type --delay %d --clearmodifiers '%s'", delayMs, text)).String()
}
