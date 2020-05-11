package Cli

import (
	"fmt"
	"strings"
)

/**
to create a beautiful box
*/
func Box(text string) {
	length := len(text)
	total := 10 + length + 10
	spaces := strings.Repeat(" ", 9)
	fmt.Println(strings.Repeat("*", total))
	fmt.Println("*" + spaces + text + spaces + "*")
	fmt.Println(strings.Repeat("*", total))
}