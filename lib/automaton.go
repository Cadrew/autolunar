package autolunar

type Automaton struct {
	state [][]uint8
	rule  *Rule
	seed  [][]uint8
}

func CreateAutomaton(rule *Rule, seed [][]uint8) *Automaton {
	state := seed
	return &Automaton{
		state: state,
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

func (am *Automaton) GetState() [][]uint8 {
	return am.state
}

func (am *Automaton) Reset() {
	am.state = am.seed
}

// Iterate iterates according to the rules
func (am *Automaton) Iterate() {
	// TODO
}

// GetStateValue converts the current state to a value
func (am * Automaton) GetStateValue() float64 {
	return 1
}