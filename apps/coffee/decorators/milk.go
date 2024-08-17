package decorators

// Decorador que adiciona leite ao café
type Milk struct {
	Decorator
}

func (m *Milk) Cost() float64 {
	return m.Beverage.Cost() + 0.30
}

func (m *Milk) Description() string {
	return m.Beverage.Description() + ", Milk"
}
