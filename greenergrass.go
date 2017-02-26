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
// with the fields separated.  If the input string is empty, then the Name will be returned with zero values appropriately.
// If the input cannot be split on sep, then Name.first will be set as the entire input string.
func SeparateName(full string, sep string) Name {

	if full == "" {
		return Name{}
	}

	var firstName string
	var lastName string

	parts := strings.Split(full, sep)

	if len(parts) == 1 {
		return Name{first: full}
	}

	commaIndex := strings.IndexAny(full, ",")
	if commaIndex != -1 {
		lastName = string(full[:commaIndex])
		firstName = parts[1]
	} else {
		lastName = parts[len(parts)-1]
		firstName = parts[0]
	}

	// If parts is > 2, then any name in between the first and last will be considered the middle name.
	var midName string
	if len(parts) > 2 {
		midName = strings.Join(parts[1:len(parts)-1], " ")
	}

	return Name{first: firstName, middle: midName, last: lastName}
}
