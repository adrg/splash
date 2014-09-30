package splash

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	Reset Property = iota
	Bold
	Dim
	Italic
	Underline
	Blink
	FastBlink
	Reverse
	Hidden
	CrossedOut
)

const (
	Black Property = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BgBlack Property = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

var attributes map[rune]Property = map[rune]Property{
	'b': Bold,
	'd': Dim,
	'i': Italic,
	'u': Underline,
	'B': Blink,
	'f': FastBlink,
	'r': Reverse,
	'h': Hidden,
	'c': CrossedOut,
}

var colors map[string]Property = map[string]Property{
	"black":   Black,
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"magenta": Magenta,
	"cyan":    Cyan,
	"white":   White,
}

type Property int

func (p Property) String() string {
	return fmt.Sprintf("\u001b[%dm", p)
}

func (p Property) Sprint(a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", p, fmt.Sprint(a...), Reset)
}

func (p Property) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", p, fmt.Sprintf(format, a...), Reset)
}

func (p Property) Sprintln(a ...interface{}) string {
	return fmt.Sprintf("%s%s%s\n", p, fmt.Sprint(a...), Reset)
}

type Style []Property

func NewStyle(props ...Property) Style {
	return props
}

func ParseStyle(style string) Style {
	props := []Property{}

	style = strings.ToLower(strings.Replace(style, " ", "", -1))
	if style == "" {
		return props
	}

	if style == "reset" {
		return append(props, Reset)
	}

	// Parse attributes
	tokens := strings.Split(style, "+")
	if len(tokens) > 1 {
		for _, attr := range tokens[1] {
			if prop, ok := attributes[attr]; ok {
				props = append(props, prop)
			}
		}
	}

	// Parse colors
	tokens = strings.Split(tokens[0], ":")
	if prop, ok := colors[tokens[0]]; ok {
		props = append(props, prop)
	}

	if len(tokens) > 1 {
		if prop, ok := colors[tokens[1]]; ok {
			// Add property as a background color
			props = append(props, prop+10)
		}
	}

	return props
}

func (s Style) String() string {
	var buffer bytes.Buffer
	for _, prop := range s {
		buffer.WriteString(prop.String())
	}

	return buffer.String()
}

func (s Style) Sprint(a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", s, fmt.Sprint(a...), Reset)
}

func (s Style) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", s, fmt.Sprintf(format, a...), Reset)
}

func (s Style) Sprintln(a ...interface{}) string {
	return fmt.Sprintf("%s%s%s\n", s, fmt.Sprint(a...), Reset)
}
