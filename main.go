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
	fmt.Println("Random generated number:", al.Rand(0, 50))
}