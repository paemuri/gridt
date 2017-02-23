package main

import (
	"fmt"
	"os"

	"github.com/Nhanderu/gridt"
	"github.com/Nhanderu/trena"
)

func main() {
	fmt.Println()
	width, _, err := trena.Size()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	grid, ok := gridt.NewWithCells(gridt.TopToBottom, "  ", os.Args[1:]...).FitIntoWidth(width)
	if !ok {
		fmt.Println("The arguments does not fit in the terminal width!")
		os.Exit(1)
	}
	fmt.Println(grid.String())
}
