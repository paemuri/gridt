# Gridt

[![Build Status][tag1img]][tag1link]
[![GoDoc][tag2img]][tag2link]
[![Go Report Card][tag3img]][tag3link]
[![codecov][tag4img]][tag4link]

Display contents as a grid in the terminal (or any place you want to write it)!
This library formats unidimensional lists (slices) as grids, with well-defined columns, suitable for fixed-width fonts, for the sake of readability.

Inspired by [ogham/rust-term-grid][1] and by the fact that I needed it for [Nhanderu/ipe][2].

### Example

**What we do not want:**

![What we do not want](./.assets/what-we-do-not-want.png "What we do no want")

**What we want:**

![What we want](./.assets/what-we-want.png "What we want")

### Usage

```go
// Just create an empty grid...
grid := gridt.New(gridt.TopToBottom, "  ")

// Verify if it fits in a determined width...
dim, ok := grid.FitIntoWidth(100)

// And get its string!
if ok {
  fmt.Print(dim.String())
}
```

```go
// But also, you can do a lot more!

// Create a grid with pre-defined cells...
grid := gridt.New(gridt.LeftToRight, "  ", "cell1", "cell2", "cell3")

// Manipulate the cells...
grid.Add("cell4", "cell5", "cell6")
grid.Insert(2, "cell2.1", "cell2.3")
grid.Delete(0)

// Base the size of the grid on the number of columns,
// instead of the number of caracters...
dim, ok := grid.FitIntoColumns(3)
```

### License

This project code is in the public domain. See the [LICENSE file][3].

[1]: https://github.com/ogham/rust-term-grid/
[2]: https://github.com/Nhanderu/ipe/
[3]: https://github.com/Nhanderu/gridt/blob/master/LICENSE

[tag1img]: https://travis-ci.org/Nhanderu/gridt.svg?branch=master
[tag1link]: https://travis-ci.org/Nhanderu/gridt
[tag2img]: https://godoc.org/gopkg.in/Nhanderu/gridt.v1?status.png
[tag2link]: https://godoc.org/gopkg.in/Nhanderu/gridt.v1
[tag3img]: https://goreportcard.com/badge/github.com/Nhanderu/gridt
[tag3link]: https://goreportcard.com/report/github.com/Nhanderu/gridt
[tag4img]: https://codecov.io/gh/Nhanderu/gridt/branch/master/graph/badge.svg
[tag4link]: https://codecov.io/gh/Nhanderu/gridt