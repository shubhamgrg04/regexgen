package main

import (
	"fmt"
	"github.com/shubhamgrg04/regxgen"
)

const pattern = "[1-9]{1,3}.[a-b]*"

func main() {
	strings, err := regxgen.Generate(pattern, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, str := range strings {
		fmt.Println(str)
	}
}