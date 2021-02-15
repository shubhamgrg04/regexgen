package main

import (
	"fmt"
	"github.com/shubhamgrg04/regxgen"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		printHelp()
		os.Exit(1)
	}
	pat := os.Args[1]
	countArg := os.Args[2]
	count, err := strconv.Atoi(countArg)
	if err != nil || count <= 0{
		fmt.Fprintln(os.Stderr, "Invalid count argument. It should be a valid positive integer")
		printHelp()
	}

	// optional config, can be used to define
	config := regxgen.Config{
		//max occurences in case of * & +
		RepetetionMax: 10,
		// optional Seed to get non-random strings
		// Seed: 5,
	}

	strings, err := regxgen.Generate(pat, count, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, str := range strings {
		fmt.Println(str)
	}
}

func printHelp() {
	fmt.Fprintln(os.Stderr, `Usage:
    $ cli ""[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{8}" 5
	92d718a0-668e-7806-a653ee68
	71217812-9a8f-8323-75934426
	f973c0b1-4da8-15bf-4b0bba3b
	d58367ca-24fb-61a7-86b572a5
	d443af78-e18e-ecc6-1937dfd5`)
}