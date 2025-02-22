/*
frde is a program that will integrate freude, and enable SSG from a
configuration file.
*/
package main

import (
	"fmt"

	"go.cpmachado.pt/freude"
)

func main() {
	fmt.Printf("Hello freude at version %s", freude.Version)
}
