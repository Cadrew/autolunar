package autolunar

import (
	"fmt"
	"time"
)

type Autolunar struct {
	prn       chan float64
	sleep     int
	automaton []*Automaton
}

var x float64 = 0

func CreateGenerator() *Autolunar {
	return &Autolunar{
		prn: make(chan float64, 1),
		sleep: 10,
		automaton: nil,
	}
}

// SetDefault sets the default configuration
func (al *Autolunar) SetDefault() error {
	fmt.Println("[autolunar] init default")
	al.sleep = 50
	al.RemoveAutomatons()
	fredkin, err := ReadRule("fredkin")
	if (err != nil) {
		return err
	}
	amoeba, err := ReadRule("amoeba")
	if (err != nil) {
		return err
	}
	al.AddAutomaton(fredkin, nil)
	al.AddAutomaton(amoeba, nil)
	return nil
}

// SetSleep allows to change the sleep value
func (al *Autolunar) SetSleep(sleep int) {
	al.sleep = sleep
}

func (al *Autolunar) AddAutomaton(rule *Rule, seed [][]uint8) {
	automaton := CreateAutomaton(rule, seed)
	al.automaton = append(al.automaton, automaton)
}

func (al *Autolunar) RemoveAutomatons() {
	al.automaton = nil
}

func (al *Autolunar) Rand(a, b int) int {
	fmt.Println("[autolunar] rand:", a, b)
	go al.Generate()
	time.Sleep(time.Duration(al.sleep) * time.Millisecond)
	prn := <-al.prn
	return int(prn) % (b - a) + a
}

func (al *Autolunar) Generate() {
	fmt.Println("[autolunar] generate")
	for {
		al.prn <- al.Iterate()
		select {
		case <- al.prn:
			// do nothing
		default:
			return
		}
	}
}

func (al *Autolunar) Iterate() float64 {
	x += 1
	return x
}