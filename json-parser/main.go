package main

import (
	"flag"
	"fmt"
	"json-parser/internal/validator"
	"strconv"
)

func main() {
	jsonString := flag.String("p", "", "json string")

	flag.Parse()

	if *jsonString == "" {
		panic("no json was provided")
	}

	fmt.Println("this is the json: " + *jsonString)

	validator := validator.NewJsonValidator(*jsonString)
	isValid := strconv.FormatBool(validator.Validate())
	fmt.Println("isValid:" + isValid)
}
