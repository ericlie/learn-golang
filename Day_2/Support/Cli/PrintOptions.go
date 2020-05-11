package Cli

import (
	"fmt"
	"golang/Day_2/Support/Data"
	"strings"
)

func PrintOptions(options map[string]Data.Option) {
	fmt.Println("COMMANDS:")
	for _, option := range options {
		if len(option.Arguments) > 0 {
			argument := strings.Join(option.Arguments, "] [")
			fmt.Printf("   - %s [%s]   %s \n", option.Name, argument, option.Description)
			continue
		}
		fmt.Printf("   - %-10s    %s \n", option.Name, option.Description)
	}
}
