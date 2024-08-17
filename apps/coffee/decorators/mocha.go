package decorators

// Decorador que adiciona mocha ao café
type Mocha struct {
	Decorator
}

func (mo *Mocha) Cost() float64 {
	return mo.Beverage.Cost() + 0.20
}

func (mo *Mocha) Description() string {
	return mo.Beverage.Description() + ", Mocha"
}
