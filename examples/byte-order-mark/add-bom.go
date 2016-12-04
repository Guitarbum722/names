package main

import (
	"io/ioutil"
	"log"
)

func main() {

	bom := []byte{239, 187, 191}
	chars := []byte("This has a BOM!")
	bom = append(bom, chars...)

	err := ioutil.WriteFile("test.txt", bom, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
