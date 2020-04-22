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

func SetColumn(n int) string {
	return fmt.Sprintf("%s[%dG", Escape, n)
}

func CursorUp(n int) string {
	return fmt.Sprintf("%s[%dA", Escape, n)
}

func CursorDown(n int) string {
	return fmt.Sprintf("%s[%dB", Escape, n)
}

func CursorForward(n int) string {
	return fmt.Sprintf("%s[%dC", Escape, n)
}

func CursorBackward(n int) string {
	return fmt.Sprintf("%s[%dD", Escape, n)
}
