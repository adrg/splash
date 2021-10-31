<h1 align="center">
    <img src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/splash/logo.png" alt="splash logo" />
</h1>

<p align="center">
    <a href="https://pkg.go.dev/github.com/adrg/splash">
        <img alt="pkg.go.dev documentation" src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white" />
    </a>
    <a href="https://opensource.org/licenses/MIT" rel="nofollow">
        <img alt="MIT license" src="https://img.shields.io/github/license/adrg/splash" />
    </a>
    <a href="https://goreportcard.com/report/github.com/adrg/splash">
        <img alt="Go report card" src="https://goreportcard.com/badge/github.com/adrg/splash" />
    </a>
</p>

Splash is a small package which allows styling terminal output.
It provides a set of types and functions to facilitate coloring and styling of
output text. It can be useful in CLI applications or logging libraries.

The core of the package is the **Property** type. A property represents either
a color (foreground or background) or a text attribute (bold, underline, etc.).
The package also defines the **Style** type which is a collection of properties.
Styles provide the ability to store a group of properties and reuse them when needed.
Both types come with the familiar [String](http://godoc.org/fmt#String),
[Sprint](http://godoc.org/fmt#Sprint), [Sprintf](http://godoc.org/fmt#Sprintf)
and [Sprintln](http://godoc.org/fmt#Sprintf) methods which are used the same as
the ones in the [fmt](http://godoc.org/fmt) package. Moreover, from a programming
standpoint, there is no difference between using a property and a style.

Full documentation can be found at https://pkg.go.dev/github.com/adrg/splash.

## Installation
```
go get github.com/adrg/splash
```
## Usage

#### Properties
```go
package main

import (
    "fmt"

    "github.com/adrg/splash"
)

func main() {
	reset := splash.Reset

	// Using text attributes
	fmt.Printf("%s%s%s ", splash.Bold, "To boldly go", reset)
	fmt.Printf("%s%s%s ", splash.Underline, "where no man", reset)
	fmt.Printf("%s%s%s\n", splash.Reverse, "has gone before.", reset)

	// Using foreground and background colors
	fmt.Printf("%s%s%s\n", splash.Red, "Roses are red", reset)
	fmt.Printf("%s%s%s\n", splash.BgGreen, "Here's something new:", reset)
	fmt.Printf("%s%s%s\n", splash.Magenta, "Violets are violet", reset)
	fmt.Printf("%s%s%s\n", splash.BgBlue, "Not freaking blue!", reset)

	// Combining colors with text attributes
	fmt.Printf("%s%s%s%s\n", splash.Bold, splash.Green, "Hint: lamp", reset)
	fmt.Printf("%s%s%s\n", splash.Red, splash.BgBlue, "Hint: famous plumbler")

	fmt.Println(reset)

	// Using property functions
	fmt.Println(splash.BgYellow.Sprint("Yellow there!"))
	fmt.Print(splash.Green.Sprintln("The Wicked Witch of The West"))
	fmt.Println(splash.Bold.Sprintf("%s%s", splash.Blue, "Don't feel blue!"))
}
```

<img alt="properties output" src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/splash/properties.png" width="100%">

#### Styles
```go
package main

import (
    "fmt"

    "github.com/adrg/splash"
)

func main() {
	// Using styles
	info := splash.NewStyle(splash.Green, splash.Bold)
	warning := splash.NewStyle(splash.Yellow)
	err := splash.NewStyle(splash.Red, splash.Bold)
	critical := splash.NewStyle(splash.Bold, splash.Yellow, splash.BgRed)

	fmt.Printf("%s%s%s\n", info, "INFO: I'm so informative", splash.Reset)
	fmt.Println(warning.Sprint("WARNING: You should not ignore me"))
	fmt.Print(err.Sprintln("ERROR: You can't say I didn't warn you"))
	fmt.Println(critical.Sprintf("%s %s\n", "CRITICAL:", "This should be good"))

	// Parsing styles
	// Format: foreground:background+attributes
	attr := splash.ParseStyle("+b")
	fmt.Println(attr.Sprint("Bold"))

	attrs := splash.ParseStyle("+bu")
	fmt.Println(attrs.Sprint("Bold, underline"))

	fg := splash.ParseStyle("yellow")
	fmt.Println(fg.Sprint("Yellow foreground"))

	bg := splash.ParseStyle(":red")
	fmt.Println(bg.Sprint("Red background"))

	fgAttr := splash.ParseStyle("green+b")
	fmt.Println(fgAttr.Sprint("Green foreground, bold"))

	bgAttr := splash.ParseStyle(":magenta+u")
	fmt.Println(bgAttr.Sprint("Magenta background, underline"))

	fgBg := splash.ParseStyle("cyan:red")
	fmt.Println(fgBg.Sprint("Cyan foreground, red background"))

	fgBgAttr := splash.ParseStyle("yellow:blue+b")
	fmt.Println(fgBgAttr.Sprint("Yellow foreground, blue background, bold"))

	fgBgAttrs := splash.ParseStyle("red:green+br")
	fmt.Println(fgBgAttrs.Sprint("Red foreground, green background, bold, reverse"))
}
```

<img alt="styles output" src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/splash/styles.png" width="100%">

## Property reference

<img alt="property reference" src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/splash/colors.png" width="100%">

**Foreground colors**
```
Black Red Green Yellow Blue Magenta Cyan White
```

**Background colors**
```
BgBlack BgRed BgGreen BgYellow BgBlue BgMagenta BgCyan BgWhite
```

**Text attributes**
```
Reset Bold Dim Italic Underline Blink FastBlink Reverse Hidden CrossedOut
```

 * Note: unfortunately not all text attributes are supported in all terminals.

## Style parsing reference

**Format**
```
foreground:background+attributes
```

**Colors**
```
black red green yellow blue magenta cyan white
```

**Text attributes**
```
b     - Bold
d     - Dim
i     - Italic
u     - Underline
B     - Blink
f     - FastBlink
r     - Reverse
h     - Hidden
c     - CrossedOut
reset - Reset
```

## References

For more information see the [ANSI escape sequences](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors)
and [Terminal colors and formatting](http://misc.flogisoft.com/bash/tip_colors_and_formatting).

## License

Copyright (c) 2014 Adrian-George Bostan.

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT).
See [LICENSE](LICENSE) for more details.
