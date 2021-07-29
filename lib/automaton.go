package autolunar

type Automaton struct {
	units []*Cellular
	rule  *Rule
	seed  [][]uint8
}

const GRID_SIZE = 100 // grid size of the automaton

func CreateAutomaton(rule *Rule, seed [][]uint8) *Automaton {
	if (rule.GetDimensions() != 2) {
		return nil
	}
	var units []*Cellular
	for i := 0; i < GRID_SIZE; i++ {		
		for j := 0; j < GRID_SIZE; j++ {
			units = append(units, CreateCell(0, j, i))
		}
	}
	for _, xy := range seed {
		index := (xy[1] * GRID_SIZE) + xy[0]
		units[index].Set(1)
	}
	return &Automaton{
		units: units,
		rule: rule,
		seed: seed,
	}
}

func (am *Automaton) SetSeed(seed [][]uint8) {
	if (len(seed) == 0) {
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

// Iterate iterates according to the rules
func (am *Automaton) Iterate() {
	if (am.GetStateValue() <= 0) {
		return
	}
	for n := range am.units {
		neighbours := am.FindNeighboursIndex(n)
		ncount := 0
		for _, neighbour := range neighbours {
			if am.units[neighbour].State() == 1 {
				ncount++
			}
		}
		if am.units[n].State() == 1 {
			if ncount < 2 || ncount > 3 {
				am.units[n].Set(0)
			}
		} else {
			if ncount == 3 {
				am.units[n].Set(1)
			}
		}
	}
}

func (am * Automaton) FindNeighboursIndex(index int) []int {
	moore := am.rule.GetNeighborhood()
	currentX, currentY := am.units[index].XY()
	var neighbours []int
	for i := 1; i <= moore; i++ {
		for n := range am.units {
			x, y := am.units[n].XY()
			if (x == currentX && y == currentY - i) {
				neighbours = append(neighbours, n)
			} else if (x == currentX && y == currentY + i) {
				neighbours = append(neighbours, n)
			} else if (x == currentX - i && y == currentY) {
				neighbours = append(neighbours, n)
			} else if (x == currentX + i && y == currentY) {
				neighbours = append(neighbours, n)
			} else if (x == currentX - i && y == currentY - i) {
				neighbours = append(neighbours, n)
			} else if (x == currentX - i && y == currentY + i) {
				neighbours = append(neighbours, n)
			} else if (x == currentX + i && y == currentY + i) {
				neighbours = append(neighbours, n)
			} else if (x == currentX + i && y == currentY - i) {
				neighbours = append(neighbours, n)
			}
		}
	}
	return neighbours
}

// GetStateValue converts the current automaton state to a single value
func (am * Automaton) GetStateValue() float64 {
	// TODO
	return 1
}