/*
frde is a program that will integrate freude, and enable SSG from a
configuration file.
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"go.cpmachado.pt/freude"
)

var Version = "0.0.0"

func init() {
	var v bool
	flag.BoolVar(&v, "version", false, "Display version and exit")
	flag.Parse()

	if v {
		fmt.Printf("frde version frde%s", Version)
		os.Exit(0)
	}
}

func main() {
	fmt.Printf("Hello freude at version %s\n", freude.Version)
}
