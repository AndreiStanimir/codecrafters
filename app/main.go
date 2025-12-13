package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for true {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		if err != nil {
			return
		}
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		words := strings.Split(strings.TrimSpace(input), " ")
		// command = strings.TrimRight(command, "\n")
		command := words[0]
		// fmt.Fprintf(os.Stderr, "You entered: %s\n", command)
		rest := words[1:]
		if command == "exit" {
			return
		}
		if command == "echo" {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(rest, " "))
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found", command)
			fmt.Fprint(os.Stdout, "\n")
		}
	}
}
