package splash_test

import "fmt"

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
	warning := splash.NewStyle(splash.Yellow)
	err := splash.NewStyle(splash.Red, splash.Bold)
	critical := splash.NewStyle(splash.Bold, splash.Yellow, splash.BgRed)
}

func ExampleParseStyle() {
	// Bold
	attr := splash.ParseStyle("+b")

	// Bold, underline
	attrs := splash.ParseStyle("+bu")

	// Yellow foreground
	fg := splash.ParseStyle("yellow")

	// Red background
	bg := splash.ParseStyle(":red")

	// Green foreground, bold
	fgAttr := splash.ParseStyle("green+b")

	// Magenta background, underline
	bgAttr := splash.ParseStyle(":magenta+u")

	// Cyan foreground, red background
	fgBg := splash.ParseStyle("cyan:red")

	// Yellow foreground, blue background, bold
	fgBgAttr := splash.ParseStyle("yellow:blue+b")

	// Red foreground, green background, bold, reverse
	fgBgAttrs := splash.ParseStyle("red:green+br")
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
