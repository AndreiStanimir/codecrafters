package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	command = strings.TrimRight(command, "\n")

	fmt.Fprintf(os.Stdout, "%s: command not found", command)
}
