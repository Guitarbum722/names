# *GreenerGrass*
_A simple package to assist with name data._

***********

### Step 0 - Installation

```sh
$ go get -u github.com/Guitarbum722/greenergrass
```

***********

### Get Started

Load default titles, prefixes and suffixes.
Initialize with a full name, then use name fields.
```go
package main

import (
	"fmt"

	"github.com/Guitarbum722/greenergrass"
)

func main() {
	greenergrass.LoadTitleData() // Load default titles, prefixes and suffixes
	n := greenergrass.New("Donald J. Trump")
	n.SeparateName(" ")
	fmt.Printf("First: %v : Mid: %v : Last: %v", n.First, n.Middle, n.Last)
}
```

or load a csv of titles by providing a file

```go
	if err := greenergrass.LoadTitleDataCSV("titles.csv"); err != nil {
		log.Fatal(err, " ::: Ah man, wasn't able to load your file.")
	}
```
***********

