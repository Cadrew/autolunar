package main

import (
	"fmt"

	autolunar "./lib"
)

func main() {
	fmt.Println("test")
	al := autolunar.GetGenerator()
	al.Init(50)
	fmt.Println(al.Rand(657, 897))
}