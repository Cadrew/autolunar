package autolunar

type Rule struct {
	dimensions int
	neighbors  int
}

func CreateRule(dim, n int) *Rule {
	return &Rule{
		dimensions: dim,
		neighbors: n,
	}
}

func (r *Rule) SetDimensions(dim int) {
	r.dimensions = dim
}

func (r *Rule) SetNeighbors(n int) {
	r.neighbors = n
}

func (r *Rule) GetDimensions() int {
	return r.dimensions
}

func (r *Rule) GetNeighbors() int {
	return r.neighbors
}