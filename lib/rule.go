package autolunar

import (
	"fmt"
	"errors"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Rule struct {
    Name          string `json:"name"`
    Model         string `json:"model"`
    Birth         []int  `json:"birth"`
    Survive       []int  `json:"survive"`
    Neighborhood  int    `json:"moore"`
    Dimensions    int    `json:"dimensions"`
    States        []int  `json:"states"`
    BxSy          string `json:"BxSy"`
}

func ReadRule(name string) (*Rule, error) {
	if (name == "") {
		return nil, errors.New(fmt.Sprintf("[autolunar] cannot create empty rule"))
	}
	jsonFile, err := os.Open(fmt.Sprintf("./rules/%s.json", name))
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var rule Rule
	json.Unmarshal(byteValue, &rule)
	return &rule, nil
}

func CreateEmptyRule(name string) *Rule {	
	return &Rule{
		Name: "",
		Model: "",
		Dimensions: 0,
		Neighborhood: 0,
		Birth: nil,
		Survive: nil,
		States: nil,
	}
}

func (r *Rule) SetDimensions(dim int) {
	r.Dimensions = dim
}

func (r *Rule) SetNeighborhood(n int) {
	r.Neighborhood = n
}

func (r *Rule) SetModel(model string) {
	r.Model = model
}

func (r *Rule) SetBirth(birth []int) {
	r.Birth = birth
}

func (r *Rule) SetSurvive(survive []int) {
	r.Survive = survive
}

func (r *Rule) SetStates(states []int) {
	r.States = states
}

func (r *Rule) GetDimensions() int {
	return r.Dimensions
}

func (r *Rule) GetNeighborhood() int {
	return r.Neighborhood
}

func (r *Rule) GetModel() string {
	return r.Model
}

func (r *Rule) GetBirth() []int {
	return r.Birth
}

func (r *Rule) GetSurvive() []int {
	return r.Survive
}

func (r *Rule) GetStates() []int {
	return r.States
}