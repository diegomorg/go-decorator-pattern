package decorators

import "fmt"

// FilterDecorator adiciona filtros ao relatório.
type FilterDecorator struct {
	ReportDecorator
}

func (f *FilterDecorator) Generate() string {
	return fmt.Sprintf("%s with filters", f.Report.Generate())
}

func (f *FilterDecorator) GetDescription() string {
	return fmt.Sprintf("%s with Filters", f.Report.GetDescription())
}
