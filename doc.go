/*
Package splash allows styling terminal output. It provides a set of types and
functions to facilitate coloring and styling of output text. It can be useful
in CLI applications or logging libraries.

The core of the package is the Property type. A property represents either a
color (foreground or background) or a text attribute (bold, underline, etc.).
The package also defines the Style type which is a collection of properties.
Styles provide the ability to store a group of properties and reuse them when needed.
Both types come with the familiar String, Sprint, Sprintf and Sprintln methods
which are used the same as the ones in the fmt package. Moreover, from a programming
standpoint, there is no difference between using a property and a style.

Usage

Text attributes:
	fmt.Printf("%s%s%s ", splash.Bold, "To boldly go", splash.Reset)
	fmt.Printf("%s%s%s ", splash.Underline, "where no man", splash.Reset)
	fmt.Printf("%s%s\n", splash.Reverse, "has gone before.", splash.Reset)

Foreground and background colors
	fmt.Printf("%s%s%s\n", splash.Red, "Roses are red", splash.Reset)
	fmt.Printf("%s%s%s\n", splash.BgGreen, "Here's something new:", splash.Reset)
	fmt.Printf("%s%s%s\n", splash.Magenta, "Violets are violet", splash.Reset)
	fmt.Printf("%s%s%s\n", splash.BgBlue, "Not freaking blue!", splash.Reset)

Combining colors with text attributes
	fmt.Printf("%s%s%s%s\n", splash.Bold, splash.Green, "Hint: lamp", splash.Reset)
	fmt.Printf("%s%s%s\n", splash.Red, splash.BgBlue, "Hint: plumbler", splash.Reset)

Property functions
	fmt.Println(splash.BgYellow.Sprint("Yellow there!"))
	fmt.Print(splash.Green.Sprintln("The Wicked Witch of The West"))
	fmt.Println(splash.Bold.Sprintf("%s%s", splash.Blue, "Don't feel blue!"))

Styles
	info := splash.NewStyle(splash.Green, splash.Bold)
	fmt.Printf("%s%s%s\n", info, "INFO: I'm so informative", splash.Reset)

	warning := splash.NewStyle(splash.Yellow)
	fmt.Println(warning.Sprint("WARNING: You should not ignore me"))

	err := splash.NewStyle(splash.Red, splash.Bold)
	fmt.Print(err.Sprintln("ERROR: You can't say I didn't warn you"))

	critical := splash.NewStyle(splash.Bold, splash.Yellow, splash.BgRed)
	fmt.Println(critical.Sprintf("%s %s\n", "CRITICAL:", "This should be good"))

Parsing styles

Format
	foreground:background+attributes

Colors
	black red green yellow blue magenta cyan white

Text attributes
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

Examples
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
*/
package splash
