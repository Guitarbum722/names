// Package greenergrass provides a utility API to assist with integration needs
// regarding data transformation, normalization, cleanup, etc.
package greenergrass

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const testVersion = 1

// Name contains a common list fields that could be combined as a person's full name
type Name struct {
	first, middle, last, prefix, suffix string
}

// init() creates a map consisting of title prefixes and suffixes that are common.
func init() {
	// ***change param to use a env extension***
	titleList, err := titleFiles("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(titleList)
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

	commaIndex := strings.IndexAny(full, ",")
	if commaIndex != -1 {
		lastName = string(full[:commaIndex])
		full = string(full[commaIndex+1:])
		full = strings.TrimLeft(full, " ")
	}

	// parts is a slice of the full input string, or the string following the first comma if provided
	parts := strings.Split(full, sep)

	if len(parts) == 1 {
		firstName = full
	} else if len(parts) >= 2 && lastName != "" {
		firstName = string(parts[0])
		midName = strings.Join(parts[1:len(parts)], " ")
	} else {
		firstName = string(parts[0])
		midName = strings.Join(parts[1:len(parts)-1], " ")
		lastName = string(parts[len(parts)-1])
	}

	return Name{first: firstName, middle: midName, last: lastName}
}

func titleFiles(filePath string) (map[string]struct{}, error) {

	if filePath == "" {
		filePath = "examples/titles.csv"
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

	titles := make(map[string]struct{})

	for _, each := range records {
		titles[each[0]] = struct{}{}
	}
	return titles, nil
}
