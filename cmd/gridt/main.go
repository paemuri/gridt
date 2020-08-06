package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Nhanderu/gridt"
	"github.com/alecthomas/kingpin"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	topToBottom = "top-to-bottom"
	leftToRight = "left-to-right"
)

var (
	args      *[]string
	file      *string
	separator *string
	direction *string
)

func init() {

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Version("2.0.0").VersionFlag.Short('v')

	file = kingpin.Flag("file", "Get values as lines from file.").Short('f').String()
	separator = kingpin.
		Flag("separator", "What separates every value column.").
		Short('s').
		Default("  ").
		String()
	direction = kingpin.
		Flag("direction", `Whether it writes from "top-to-bottom" or "left-to-right".`).
		Short('d').
		Default(topToBottom).
		Enum(topToBottom, leftToRight)

	kingpin.Parse()
}

func main() {

	width, _, err := terminal.GetSize(0)
	if err != nil {
		eprintln(err.Error())
		os.Exit(1)
	}

	var scanner *bufio.Scanner
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			eprintf("Error %s.", err.Error())
			os.Exit(1)
		}
		defer f.Close()
		scanner = bufio.NewScanner(f)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	var values []string
	if scanner != nil {
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			values = append(values, scanner.Text())
		}
	}

	if len(values) == 0 {
		eprintln("No values were given.")
		os.Exit(1)
	}

	d := gridt.TopToBottom
	if *direction == leftToRight {
		d = gridt.LeftToRight
	}

	grid, ok := gridt.New(d, *separator, values...).FitIntoWidth(width)
	if !ok {
		eprintln("The given values does not fit in the terminal width.")
		os.Exit(1)
	}

	fmt.Println(grid.String())
}

func eprintln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}

func eprintf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}
