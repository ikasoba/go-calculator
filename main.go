package main

import (
	"fmt"

	"bufio"
	"os"

	"github.com/mattn/go-isatty"
)

func main() {
	isTerminal := isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
	scanner := bufio.NewScanner(os.Stdin)

	runtime := &Runtime{}

	for {
		if isTerminal {
			fmt.Print("> ")
		}

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		tree, _, err := ParseExpr(0, line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		value, err := runtime.Exec(tree)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(value)
	}
}
