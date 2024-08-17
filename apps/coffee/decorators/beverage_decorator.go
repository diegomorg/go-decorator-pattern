package decorators

import "diegomorg/go-decorator/apps/coffee/interfaces"

// Estrutura base que todos os decoradores herdam
type Decorator struct {
	Beverage interfaces.Beverage
}

func (d *Decorator) Cost() float64 {
	return d.Beverage.Cost()
}

func (d *Decorator) Description() string {
	return d.Beverage.Description()
}
