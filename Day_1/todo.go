package main

import (
	"../src/Struct/Command.go"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	box("Todo and GO! - By Eric")

	commands := map[string]string{
		"h": "Shows help information",
		"a": "Add new Task",
		"l": "List all tasks",
	}


	listCommands(commands)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if text == "x" {
			box("Thank you!")
			os.Exit(0)
		}
		echo("your input is: " + text)
		if strings.Compare("hi", text) == 0 {
			echo("hello, Yourself")
		}
	}
}

func echo(text string) {
	fmt.Println(text)
}

/**
to create a beautiful box
*/
func box(text string) {
	length := len(text)
	total := 10 + length + 10
	spaces := strings.Repeat(" ", 9)
	echo(strings.Repeat("*", total))
	echo("*" + spaces + text + spaces + "*")
	echo(strings.Repeat("*", total))
}

func listCommands(commands map[string]string) {
	echo("COMMANDS:")
	for command, intro := range commands {
		fmt.Printf("* %s - %s \n", command, intro)
	}
}
