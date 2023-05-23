// This program takes language tags as command-line
// arguments and parses them.

package main

import (
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

func main() {
	for _, arg := range os.Args[1:] {
		tag, err := language.Parse(arg)
		if err != nil {
			fmt.Printf("%s: error: %v\n", arg, err)
		} else if tag == language.Und {
			fmt.Printf("%s: undefined\n", arg)
		} else {
			fmt.Printf("%s: tag %s\n", arg, tag)
		}
	}
	json := `{"hello": "world"}`
	if gjson.Valid(json) {
		r := gjson.Get(json, "hello")
		fmt.Printf("hello: %v\n", r)
	}
}
