package main

import (
	"bufio"
	"fmt"
	"golang/Day_2/Support/Cli"
	"golang/Day_2/Support/Data"
	"os"
	"strings"
)

var tasks map[string]Data.Task

func main() {
	tasks = make(map[string]Data.Task)
	options := map[string]Data.Option{
		"Add": {
			Name:        "Add",
			Description: "Add new Task. Usage: add \" Learn GoLang! 万岁! \"",
			Arguments: []string{
				"Task",
			},
			Action: func(input string) {
				fmt.Println("Add: Your input is==> " + input)
				inputs := strings.Split(input, " ")
				tasks[inputs[1]] = Data.Task{
					Id:    len(tasks) + 1,
					Title: inputs[1],
				}
			},
		},
		"Help": {
			Name:        "Help",
			Description: "Show help",
			Action: func(input string) {
				fmt.Println("HELP input is==> " + input)
			},
		},
		"List": {
			Name:        "List",
			Description: "Show all tasks",
			Action: func(input string) {
				for title, task := range tasks {
					fmt.Printf("   - %d %-10s    %t \n", task.Id, title, task.IsDone)
				}
			},
		},
	}

	reader := bufio.NewReader(os.Stdin)

	Cli.Box("Hello World!")
	Cli.PrintOptions(options)

	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println("Your input: " + input)
		inputs := strings.Split(input, " ")
		var opt Data.Option
		if len(inputs) > 1 {
			opt = options[inputs[0]]
		} else {
			opt = options[input]
		}
		fmt.Println(opt.Description)
		opt.Action(input)
	}
}
