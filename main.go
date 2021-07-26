package main

import (
	"fmt"

	autolunar "./lib"
)

// usage example
func main() {
	al := autolunar.GetGenerator()
	al.SetDefault()
	fmt.Println(al.Rand(0, 897))
}