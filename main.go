package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"

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
	// generate a number
	RNG := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		RNG[i] = al.Rand(0, 100000)
	}
	fmt.Println("Random generated number:", RNG)
 
	file, err := os.OpenFile("numbers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 
	if err != nil {
		fmt.Printf("failed creating file: %s", err)
	}
 
	datawriter := bufio.NewWriter(file)
 
	for _, data := range RNG {
		_, _ = datawriter.WriteString(strconv.Itoa(data) + "\n")
	}
 
	datawriter.Flush()
	file.Close()
}