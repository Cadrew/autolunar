package autolunar

type Automaton struct {
	state chan int
	rule  *Rule
	seed  [][]uint8
}

func CreateAutomaton(rule *Rule, seed [][]uint8) *Automaton {
	return &Automaton{
		state: make(chan int, 1),
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

func (am *Automaton) GetState() int {
	select {
	case state := <-am.state:
		return state
	default:
		return -1
	}
}