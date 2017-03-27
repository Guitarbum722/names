# *Names*

[![Go Report Card](https://goreportcard.com/badge/github.com/Guitarbum722/names)](https://goreportcard.com/report/github.com/Guitarbum722/names)
[![Coverage Status](https://img.shields.io/badge/coverage-93.9%25-brightgreen.svg?style=flat-square)](http://gocover.io/github.com/Guitarbum722/names)
[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/Guitarbum722/names) 


_A simple package to assist with name data._

***********

### Step 0 - Installation

```sh
$ go get -u github.com/Guitarbum722/names
```

***********

### Get Started

Load default titles, prefixes and suffixes.
Initialize with a full name, then use name fields.
```go
package main

import (
	"fmt"

	"github.com/Guitarbum722/names"
)

func main() {
	names.LoadTitleData() // Load default titles, prefixes and suffixes
	n := names.New("Donald J. Trump")
	n.SeparateName(" ")
	fmt.Printf("First: %v : Mid: %v : Last: %v", n.First, n.Middle, n.Last)
}
```

...or pass a bytes.Buffer from a CSV of your choice

```go
	if err := names.LoadTitleDataCSV(bytes.NewBuffer([]byte("Mrs.,p\nMr.,p")); err != nil {
		log.Fatal(err, " ::: Ah man, wasn't able to load your file.")
	}
```

Get the initials for the name.

```go
    // include periods
    fmt.Println(n.Initials(true)) // D.J.T.

    // exclude periods
    fmt.Println(n.Initials(false)) // DJT
```
***********

