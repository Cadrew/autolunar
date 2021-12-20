package autolunar

import (
	"fmt"
	"time"
)

type Autolunar struct {
	prn           chan float64
	sleep         int
	automaton     []*Automaton
	previousRound int
}

const ROUND = 100

// CreateGenerator return the instance of a new generator
func CreateGenerator() *Autolunar {
	return &Autolunar{
		prn:           make(chan float64, 1),
		sleep:         10,
		automaton:     nil,
		previousRound: 0,
	}
}

// SetDefault sets the default configuration
func (al *Autolunar) SetDefault() error {
	fmt.Println("[autolunar] init default")
	al.sleep = 10
	al.RemoveAutomata()
	fredkin, err := ReadRule("fredkin")
	if err != nil {
		return err
	}
	amoeba, err := ReadRule("amoeba")
	if err != nil {
		return err
	}
	gun, err := ReadSeed("gun")
	if err != nil {
		return err
	}
	al.AddAutomaton(fredkin, gun)
	al.AddAutomaton(amoeba, [][]uint8{
		{1, 5}, {1, 6}, {2, 5}, {2, 6},
		{11, 5}, {11, 6}, {11, 7}, {12, 4}, {12, 8}, {13, 3},
		{13, 9}, {14, 3}, {14, 9}, {15, 6}, {16, 4}, {16, 8},
		{21, 3}, {21, 4}, {21, 5}, {22, 3}, {22, 4}, {22, 5},
		{23, 2}, {23, 6}, {25, 1}, {25, 2},
		{35, 3}, {35, 4}, {36, 3}, {36, 4},
	})
	return nil
}

// SetSleep allows to change the sleep value
func (al *Autolunar) SetSleep(sleep int) {
	al.sleep = sleep
}

// AddAutomaton adds a new automaton in the generator (needs to provide the rule)
func (al *Autolunar) AddAutomaton(rule *Rule, seed Seed) {
	automaton := CreateAutomaton(rule, seed)
	al.automaton = append(al.automaton, automaton)
}

// RemoveAutomata removes all the automata in our generator
func (al *Autolunar) RemoveAutomata() {
	al.automaton = nil
}

// Rand returns the pseudo random number
func (al *Autolunar) Rand(a, b int) int {
	go al.Generate()
	time.Sleep(time.Duration(al.sleep) * time.Millisecond)
	prn := <-al.prn
	al.previousRound = int((prn - float64(int64(prn))) * ROUND)
	return int(prn+float64(getTimestamp()))%(b-a) + a
}

// Generate iterates into the automata until the al.prn channel is read
func (al *Autolunar) Generate() {
	for {
		automaton := al.previousRound % len(al.automaton)
		al.automaton[automaton].Iterate()
		al.prn <- al.automaton[automaton].GetStateValue()
		select {
		case prn := <-al.prn:
			al.previousRound = int((prn - float64(int64(prn))) * ROUND)
		default:
			return
		}
	}
}
