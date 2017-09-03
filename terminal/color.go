package terminal

import "fmt"

//Yellow add ANSI escape codes to the given string so when printed on a terminal
//the text foreground is yellow.
func Yellow(s string) string {
	return colorize(s, 220)
}

//Red add ANSI escape codes to the given string so when printed on a terminal
//the text foreground is red.
func Red(s string) string {
	return colorize(s, 88)
}

//White add ANSI escape codes to the given string so when printed on a terminal
//the text foreground is white.
func White(s string) string {
	return colorize(s, 244)
}

//colorize applies the given color to the given text using the "extended set foreground color" code
func colorize(s string, color uint8) string {
	return fmt.Sprintf("\033[38;5;%03dm%s\033[0m", color, s)
}
