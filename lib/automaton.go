package autolunar

import (
	"fmt"
	"math"
)

type Automaton struct {
	units []*Cellular
	rule  *Rule
	seed  [][]uint8
}

// grid size of the automaton 64x64 (2^8 x 2^8)
// it will be easier to do block of bytes
const GRID_SIZE = 64

func CreateAutomaton(rule *Rule, seed [][]uint8) *Automaton {
	if rule.GetDimensions() != 2 {
		return nil
	}
	var units []*Cellular
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			units = append(units, CreateCell(0, j, i))
		}
	}
	for _, xy := range seed {
		var index int64 = (int64(xy[1]) * int64(GRID_SIZE)) + int64(xy[0])
		units[index].Set(1)
	}
	return &Automaton{
		units: units,
		rule:  rule,
		seed:  seed,
	}
}

func (am *Automaton) SetSeed(seed [][]uint8) {
	if len(seed) == 0 {
		return
	}
	am.seed = seed
}

func (am *Automaton) SetRule(rule *Rule) {
	am.rule = rule
}

func (am *Automaton) GetSeed() [][]uint8 {
	return am.seed
}

func (am *Automaton) GetRule() *Rule {
	return am.rule
}

func (am *Automaton) GetUnits() []*Cellular {
	return am.units
}

func (am *Automaton) Reset() {
	var units []*Cellular
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			units = append(units, CreateCell(0, j, i))
		}
	}
	for _, xy := range am.seed {
		index := (xy[1] * GRID_SIZE) + xy[0]
		units[index].Set(1)
	}
	am.units = units
}

// Iterate iterates accordingly to the rules
func (am *Automaton) Iterate() {
	new_units := am.units
	for n := range am.units {
		neighbours := am.FindNeighboursIndex(n)
		ncount := 0
		for _, neighbour := range neighbours {
			if am.units[neighbour].State() == 1 {
				ncount++
			}
		}
		if am.units[n].State() == 1 {
			// death
			if !contains(am.rule.GetSurvive(), ncount) {
				new_units[n].Set(0)
			}
		} else {
			// birth
			if contains(am.rule.GetBirth(), ncount) {
				new_units[n].Set(1)
			}
		}
	}
	am.units = new_units
}

// FindNeighboursIndex returns the index of all the neighbours depending on moore neighborhood
func (am *Automaton) FindNeighboursIndex(index int) []int {
	moore := am.rule.GetNeighborhood()
	currentX, currentY := am.units[index].XY()
	var neighbours []int
	for i := 1; i <= moore; i++ {
		if currentY-i < GRID_SIZE && currentX < GRID_SIZE && currentY-i >= 0 && currentX >= 0 {
			neighbours = append(neighbours, ((currentY-i)*GRID_SIZE)+currentX)
		}
		if currentY+i < GRID_SIZE && currentX < GRID_SIZE && currentY+i >= 0 && currentX >= 0 {
			neighbours = append(neighbours, ((currentY+i)*GRID_SIZE)+currentX)
		}
		if currentY < GRID_SIZE && currentX-i < GRID_SIZE && currentY >= 0 && currentX-i >= 0 {
			neighbours = append(neighbours, (currentY*GRID_SIZE)+currentX-i)
		}
		if currentY < GRID_SIZE && currentX+i < GRID_SIZE && currentY >= 0 && currentX+i >= 0 {
			neighbours = append(neighbours, (currentY*GRID_SIZE)+currentX+i)
		}
		if currentY-i < GRID_SIZE && currentX-i < GRID_SIZE && currentY-i >= 0 && currentX-i >= 0 {
			neighbours = append(neighbours, ((currentY-i)*GRID_SIZE)+currentX-i)
		}
		if currentY+i < GRID_SIZE && currentX-i < GRID_SIZE && currentY+i >= 0 && currentX-i >= 0 {
			neighbours = append(neighbours, ((currentY+i)*GRID_SIZE)+currentX-i)
		}
		if currentY+i < GRID_SIZE && currentX+i < GRID_SIZE && currentY+i >= 0 && currentX+i >= 0 {
			neighbours = append(neighbours, ((currentY+i)*GRID_SIZE)+currentX+i)
		}
		if currentY-i < GRID_SIZE && currentX+i < GRID_SIZE && currentY-i >= 0 && currentX+i >= 0 {
			neighbours = append(neighbours, ((currentY-i)*GRID_SIZE)+currentX+i)
		}
	}

	return neighbours
}

// GetStateValue converts the current automaton state to a single value
func (am *Automaton) GetStateValue() float64 {
	var value float64 = 0
	for n := range am.units {
		value += float64(am.units[n].State()) * float64(n)
	}
	return float64(value) / ROUND
}

// GetStateValue calculates a value from block of bytes
func (am *Automaton) GetBytesValue() float64 {
	blocks := make([]int64, 8)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			blocks[i] += int64(math.Pow(float64(2), float64(8*j))) * int64(am.units[8*i+j].State())
		}
	}
	// vector = blocks[0] ??? blocks[1] ??? . . . ??? blocks[7]
	vector := int64(0)
	for i := 0; i < 8; i++ {
		vector ^= blocks[i]
	}
	return float64(vector) / ROUND
}

func (am *Automaton) Display() {
	for i := 0; i < GRID_SIZE*GRID_SIZE; i += GRID_SIZE {
		for j := 0; j < GRID_SIZE; j++ {
			fmt.Print(am.units[i+j].State())
		}
		fmt.Println()
	}
}
