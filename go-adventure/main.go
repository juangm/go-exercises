package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/juangm/go-exercises/go-adventure/adventure"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the adventure story")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		// Not a good idea...
		panic(err)
	}

	d := json.NewDecoder(f)
	var story adventure.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", story)
}