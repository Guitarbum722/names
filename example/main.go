package main

import (
	"github.com/tidwall/gjson"
)

func main() {
	jsonSample := `{ "name": {"first": "John", "last": "Mo世界"}}`

	value := gjson.Get(jsonSample, "name.last")
	println(value.String())
}
