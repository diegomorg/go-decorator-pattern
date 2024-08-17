package decorators

import "fmt"

// ExportDecorator adiciona opções de exportação ao relatório.
type ExportDecorator struct {
	ReportDecorator
}

func (e *ExportDecorator) Generate() string {
	return fmt.Sprintf("%s with export options", e.Report.Generate())
}

func (e *ExportDecorator) GetDescription() string {
	return fmt.Sprintf("%s with Export Options", e.Report.GetDescription())
}
