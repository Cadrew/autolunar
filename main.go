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
	// generate a number between 0 and 50
	fmt.Println("Random generated number:", al.Rand(0, 50))
}