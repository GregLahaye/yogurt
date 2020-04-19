package yogurt

import "fmt"

const Escape = "\x1B"
const ResetForeground = Escape + "[39m"
const ResetBackground = Escape + "[49m"
const EnableBlinking = Escape + "[?12h"
const DisableBlinking = Escape + "[?12l"
const EnableCursor = Escape + "[?25h"
const DisableCursor = Escape + "[?25l"
const ClearLine = Escape + "[2K"

func Foreground(color string) string {
	return fmt.Sprintf("%s[38;5;%sm", Escape, color)
}

func Background(color string) string {
	return fmt.Sprintf("%s[48;5;%sm", Escape, color)
}

func SetCursor(line, col int) {
	fmt.Printf("%s[%d;%dH", Escape, line, col)
}

func SetColumn(col int) {
	fmt.Printf("%s[%dG", Escape, col)
}

func CursorUp(value int) {
	fmt.Printf("%s[%dA", Escape, value)
}

func CursorDown(value int) {
	fmt.Printf("%s[%dB", Escape, value)
}

func CursorForward(value int) {
	fmt.Printf("%s[%dC", Escape, value)
}

func CursorBackward(value int) {
	fmt.Printf("%s[%dD", Escape, value)
}
