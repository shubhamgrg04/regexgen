package main

import (
	"fmt"
	"github.com/shubhamgrg04/regxgen"
)

//const pattern = "[1-9]{1,3}.[a-b]*"
//const pattern = "[-+]?[0-9]{1,16}[.][0-9]{1,6}"
//const pattern = "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{8}"
const pattern = "[^aeiouAEIOU0-9]{5}"
func main() {
	// optional config, can be used to define
	config := regxgen.Config{
		//max occurences in case of * & +
		RepetetionMax: 10,
		// optional Seed to get non-random strings
		// Seed: 5,
	}
	strings, err := regxgen.Generate(pattern, 200, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, str := range strings {
		fmt.Println(str)
	}
}