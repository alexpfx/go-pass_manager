package xdotool

import (
	"fmt"

	"github.com/bitfield/script"
)

func (x Xdotool) Type(text string, delayMs int) (string, error) {
	return script.Exec(fmt.Sprintf("xdotool type --delay %d --clearmodifiers '%s'", delayMs, text)).String()
}

type Xdotool struct{}
