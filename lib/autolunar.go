package autolunar

import (
	"fmt"
	"time"
)

type Autolunar struct {
	PRN chan float64
	sleep int
}

// var PRN = make(chan float64)
var x float64 = 0

func GetGenerator() *Autolunar {
	return &Autolunar{
		PRN: make(chan float64, 1),
		sleep: 10,
	}
}

// Init uses CA rules to init the generator
func (al *Autolunar) Init(sleep int) {
	fmt.Println("[autolunar] init")
	al.sleep = sleep
}

func (al *Autolunar) Rand(a, b int) int {
	fmt.Println("[autolunar] rand:", a, b)
	go al.Generate()
	time.Sleep(time.Duration(al.sleep) * time.Millisecond)
	prn := <-al.PRN
	return int(prn) % (b - a) + a
}

func (al *Autolunar) Generate() {
	fmt.Println("[autolunar] generate")
	for {
		al.PRN <- al.Iterate()
		<- al.PRN
	}
}

func (al *Autolunar) Iterate() float64 {
	fmt.Println("[autolunar] iterate:", x)
	x += 1
	return x
}