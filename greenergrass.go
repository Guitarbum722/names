// Package greenergrass provides a utility API to assist with integration needs
// regarding data transformation, normalization, cleanup, etc.
package greenergrass

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const testVersion = 1

// Name contains a common list fields that could be combined as a person's full name
type Name struct {
	first, middle, last, prefix, suffix string
}

// LoadTitleData creates a map consisting of title prefixes and suffixes that are common.
// This can be called optionally by the consumer if they are expecting their input data to
// include prefixes and/or suffixes
func LoadTitleData() error {
	// ***change param to use a env extension***
	_, err := titleFiles("")
	if err != nil {
		return err
	}
	return nil
}

// SeparateName accepts a name n as input, and parses it according to common logic and returns a Name struct
// with the fields separated.  If the input string is empty, then the Name will be returned with zero values appropriately.
// If the input cannot be split on sep, then Name.first will be set as the entire input string.
func SeparateName(full string, sep string) Name {
	if full == "" {
		return Name{}
	}
	if sep == "" {
		sep = " "
	}

	var firstName string
	var lastName string
	var midName string
	var pref string
	var suff string

	commaIndex := strings.IndexAny(full, ",")
	if commaIndex != -1 {
		lastName = string(full[:commaIndex])
		full = string(full[commaIndex+1:])
		full = strings.TrimLeft(full, " ")
	}

	// parts is a slice of the full input string, or the string following the first comma if provided
	parts := strings.Split(full, sep)

	// check titleList to see if the first word of full is a listed prefix
	if _, ok := titleList[parts[0]]; ok {
		pref = parts[0]
		parts = parts[1:]
	}

	// check titleList to see if the last word of full is a listed suffix or title
	if _, ok := titleList[parts[len(parts)-1]]; ok {
		suff = parts[len(parts)-1]
		parts = parts[:len(parts)-1]
	}

	if len(parts) == 1 {
		firstName = parts[0]
	} else if len(parts) >= 2 && lastName != "" {
		firstName = string(parts[0])
		midName = strings.Join(parts[1:len(parts)], " ")
	} else {
		firstName = string(parts[0])
		midName = strings.Join(parts[1:len(parts)-1], " ")
		lastName = string(parts[len(parts)-1])
	}

	return Name{first: firstName, middle: midName, last: lastName, prefix: pref, suffix: suff}
}

var titleList = make(map[string]struct{})

func titleFiles(filePath string) (map[string]struct{}, error) {

	if filePath == "" {
		filePath = "titles.csv"
	}

	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "error opening csv")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "error reading csv")
	}

	for _, each := range records {
		titleList[each[0]] = struct{}{}
	}
	return titleList, nil
}
