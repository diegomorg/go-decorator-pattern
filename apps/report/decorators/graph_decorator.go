package decorators

import "fmt"

// GraphDecorator adiciona gráficos ao relatório.
type GraphDecorator struct {
	ReportDecorator
}

func (g *GraphDecorator) Generate() string {
	return fmt.Sprintf("%s with graphs", g.Report.Generate())
}

func (g *GraphDecorator) GetDescription() string {
	return fmt.Sprintf("%s with Graphs", g.Report.GetDescription())
}
