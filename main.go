package main

import (
	"fmt"

	autolunar "./lib"
)

// usage example
func main() {
	al := autolunar.CreateGenerator()
	err := al.SetDefault()
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(al.Rand(0, 897))
}