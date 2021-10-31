package splash_test

import (
	"fmt"

	"github.com/adrg/splash"
)

func ExampleProperty_Sprint() {
	fmt.Println(splash.Yellow.Sprint("Yellow there!"))
}

func ExampleProperty_Sprintln() {
	fmt.Print(splash.Green.Sprintln("The Wicked Witch of The West"))
}

func ExampleProperty_Sprintf() {
	fmt.Println(splash.Bold.Sprintf("%s%s", splash.Blue, "Don't feel blue!"))
}

func ExampleNewStyle() {
	info := splash.NewStyle(splash.Green, splash.Bold)
	fmt.Println(info.Sprint("INFO: I'm so informative"))

	warning := splash.NewStyle(splash.Yellow)
	fmt.Println(warning.Sprint("WARNING: You should not ignore me"))

	err := splash.NewStyle(splash.Red, splash.Bold)
	fmt.Println(err.Sprint("ERROR: You can't say I didn't warn you"))

	critical := splash.NewStyle(splash.Bold, splash.Yellow, splash.BgRed)
	fmt.Println(critical.Sprint("ERROR: This should be good"))
}

func ExampleParseStyle() {
	// Bold.
	attr := splash.ParseStyle("+b")
	fmt.Println(attr.Sprint("Bold"))

	// Bold, underline.
	attrs := splash.ParseStyle("+bu")
	fmt.Println(attrs.Sprint("Bold, underline"))

	// Yellow foreground.
	fg := splash.ParseStyle("yellow")
	fmt.Println(fg.Sprint("Yellow foreground"))

	// Red background.
	bg := splash.ParseStyle(":red")
	fmt.Println(bg.Sprint("Red background"))

	// Green foreground, bold.
	fgAttr := splash.ParseStyle("green+b")
	fmt.Println(fgAttr.Sprint("Green foreground, bold"))

	// Magenta background, underline.
	bgAttr := splash.ParseStyle(":magenta+u")
	fmt.Println(bgAttr.Sprint("Magenta background, underline"))

	// Cyan foreground, red background.
	fgBg := splash.ParseStyle("cyan:red")
	fmt.Println(fgBg.Sprint("Cyan foreground, red background"))

	// Yellow foreground, blue background, bold.
	fgBgAttr := splash.ParseStyle("yellow:blue+b")
	fmt.Println(fgBgAttr.Sprint("Yellow foreground, blue background, bold"))

	// Red foreground, green background, bold, reverse.
	fgBgAttrs := splash.ParseStyle("red:green+br")
	fmt.Println(fgBgAttrs.Sprint("Red foreground, green background, bold, reverse"))
}

func ExampleStyle_Sprint() {
	warning := splash.NewStyle(splash.Yellow)
	fmt.Println(warning.Sprint("WARNING: You should not ignore me"))
}

func ExampleStyle_Sprintln() {
	err := splash.NewStyle(splash.Red, splash.Bold)
	fmt.Print(err.Sprintln("ERROR: You can't say I didn't warn you"))
}

func ExampleStyle_Sprintf() {
	critical := splash.NewStyle(splash.Bold, splash.Yellow, splash.BgRed)
	fmt.Println(critical.Sprintf("%s %s\n", "CRITICAL:", "This should be good"))
}
