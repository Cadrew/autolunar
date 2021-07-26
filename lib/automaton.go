package autolunar

type Automaton struct {
	state chan int
	rule  *Rule
	seed  int
}

func CreateAutomaton(rule *Rule, seed int) *Automaton {
	return &Automaton{
		state: make(chan int, 1),
		rule: rule,
		seed: seed,
	}
}

func (am *Automaton) SetSeed(seed int) {
	if (seed < 0) {
		return
	}
	am.seed = seed
}

func (am *Automaton) SetRule(rule *Rule) {
	am.rule = rule
}

func (am *Automaton) GetSeed() int {
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