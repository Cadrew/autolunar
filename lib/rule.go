package autolunar

import (
	"fmt"
	"errors"
)

type Rule struct {
	model      string
	dimensions int
	neighbors  int
	birth      []int
	survive    []int
	states     []int
}

func ReadRule(name string) (*Rule, error) {
	if (name == "") {
		return nil, errors.New(fmt.Sprintf("[autolunar] cannot create rule %s", name))
	}
	
	return &Rule{
		model: "",
		dimensions: 0,
		neighbors: 0,
		birth: nil,
		survive: nil,
		states: nil,
	}, nil
}

func CreateEmptyRule(name string) *Rule {	
	return &Rule{
		model: "",
		dimensions: 0,
		neighbors: 0,
		birth: nil,
		survive: nil,
		states: nil,
	}
}

func (r *Rule) SetDimensions(dim int) {
	r.dimensions = dim
}

func (r *Rule) SetNeighbors(n int) {
	r.neighbors = n
}

func (r *Rule) SetModel(model string) {
	r.model = model
}

func (r *Rule) SetBirth(birth []int) {
	r.birth = birth
}

func (r *Rule) SetSurvive(survive []int) {
	r.survive = survive
}

func (r *Rule) SetStates(states []int) {
	r.states = states
}

func (r *Rule) GetDimensions() int {
	return r.dimensions
}

func (r *Rule) GetNeighbors() int {
	return r.neighbors
}

func (r *Rule) GetModel() string {
	return r.model
}

func (r *Rule) GetBirth() []int {
	return r.birth
}

func (r *Rule) GetSurvive() []int {
	return r.survive
}

func (r *Rule) GetStates() []int {
	return r.states
}