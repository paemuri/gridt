package main

import (
	"fmt"
	"os"

	"github.com/Nhanderu/gridt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	fmt.Println()
	width, _, err := terminal.GetSize(0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	grid, ok := gridt.New(gridt.TopToBottom, "  ", os.Args[1:]...).FitIntoWidth(width)
	if !ok {
		fmt.Println("The arguments does not fit in the terminal width!")
		os.Exit(1)
	}

	fmt.Println(grid.String())
}
