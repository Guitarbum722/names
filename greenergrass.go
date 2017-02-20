// Package greenergrass provides a utility API to assist with integration needs
// regarding data transformation, normalization, cleanup, etc.
package greenergrass

import "strings"

const testVersion = 1

// Name contains a common list fields that could be combined as a person's full name
type Name struct {
	first  string
	middle string
	last   string
	prefix string
	suffix string
}

// SeparateName accepts a name n as input, and parses it according to common logic and returns a Name struct
// with the fields separated.
func SeparateName(n string) Name {
	parts := strings.Split(n, " ")
	var midName = ""
	if len(parts) > 2 {
		midName = strings.Join(parts[1:len(parts)-1], " ")
	}
	return Name{first: parts[0], middle: midName, last: parts[len(parts)-1]}
}
