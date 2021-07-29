package autolunar

type Cellular struct {
	state int
	x     int
	y     int
}

func CreateCell(state, x, y int) *Cellular {
	return &Cellular {
		state: state,
		x: x,
		y: y,
	}
}

func (c *Cellular) XY() (int, int) {
	return c.x, c.y
}

func (c *Cellular) State() int {
	return c.state
}

func (c *Cellular) Set(state int) {
	c.state = state
}