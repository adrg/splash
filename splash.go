package splash

import (
	"bytes"
	"fmt"
	"strings"
)

// Attributes.
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

// Foreground colors.
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

// Background colors.
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

var attributes = map[rune]Property{
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

var colors = map[string]Property{
	"black":   Black,
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"magenta": Magenta,
	"cyan":    Cyan,
	"white":   White,
}

// Property represents either a color (foreground or background) or a text
// attribute (bold, underline, etc.).
type Property int

// String returns the printable string representation of the property.
func (p Property) String() string {
	return fmt.Sprintf("\u001b[%dm", p)
}

// Sprint formats using the default formats for its operands and returns the
// resulting string. Spaces are added between operands when neither is a
// string. The output string is wrapped in the value of the property and
// a reset is applied at end.
func (p Property) Sprint(a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", p, fmt.Sprint(a...), Reset)
}

// Sprintf formats according to a format specifier and returns the
// resulting string. The output string is wrapped in the value of the property
// and a style is applied at end.
func (p Property) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", p, fmt.Sprintf(format, a...), Reset)
}

// Sprintln formats using the default formats for its operands and returns
// the resulting string. Spaces are always added between operands and a
// newline is appended. The output string is wrapped in the value of the
// property and a reset is applied at end.
func (p Property) Sprintln(a ...interface{}) string {
	return fmt.Sprintf("%s%s%s\n", p, fmt.Sprint(a...), Reset)
}

// PromptEscape encloses the property to be displayed in the shell prompt
// string in the escape sequences \[ and \].
func (p Property) PromptEscape() string {
	return fmt.Sprintf(`\[%s\]`, p)
}

// PromptString styles the input string, escaping the property in order for it
// to be displayed in the shell prompt string. A prompt escaped style reset is
// also applied at the end of the output string.
func (p Property) PromptString(str string) string {
	return fmt.Sprintf("%s%s%s", p.PromptEscape(), str, Reset.PromptEscape())
}

// Style is a reusable collection of properties.
type Style []Property

// NewStyle returns a new style from property arguments passed in.
func NewStyle(props ...Property) Style {
	return props
}

// ParseStyle returns a new style from the properties specified in the style
// parameter.
func ParseStyle(style string) Style {
	props := []Property{}

	style = strings.TrimSpace(strings.Replace(style, " ", "", -1))
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
	tokens = strings.Split(strings.ToLower(tokens[0]), ":")
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

// String returns the printable string representation of the style.
func (s Style) String() string {
	var buffer bytes.Buffer
	for _, prop := range s {
		buffer.WriteString(prop.String())
	}

	return buffer.String()
}

// Sprint formats using the default formats for its operands and returns the
// resulting string. Spaces are added between operands when neither is a
// string. The output string is wrapped in the value of the style and a reset
// is applied at end.
func (s Style) Sprint(a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", s, fmt.Sprint(a...), Reset)
}

// Sprintf formats according to a format specifier and returns the
// resulting string. The output string is wrapped in the value of the style
// and a reset is applied at end.
func (s Style) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", s, fmt.Sprintf(format, a...), Reset)
}

// Sprintln formats using the default formats for its operands and returns
// the resulting string. Spaces are always added between operands and a
// newline is appended. The output string is wrapped in the value of the style
// and a reset is applied at end.
func (s Style) Sprintln(a ...interface{}) string {
	return fmt.Sprintf("%s%s%s\n", s, fmt.Sprint(a...), Reset)
}

// PromptEscape encloses the style to be displayed in the shell prompt
// string in the escape sequences \[ and \].
func (s Style) PromptEscape() string {
	return fmt.Sprintf(`\[%s\]`, s)
}

// PromptString styles the input string, escaping the style in order for it to
// be displayed in the shell prompt string. A prompt escaped style reset is
// also applied at the end of the output string.
func (s Style) PromptString(str string) string {
	return fmt.Sprintf("%s%s%s", s.PromptEscape(), str, Reset.PromptEscape())
}
