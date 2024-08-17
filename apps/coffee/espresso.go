package coffee

// Implementação concreta da interface Beverage
type Espresso struct{}

func (e *Espresso) Cost() float64 {
	return 1.99
}

func (e *Espresso) Description() string {
	return "Espresso"
}
