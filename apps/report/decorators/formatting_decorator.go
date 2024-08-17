package decorators

import "fmt"

// FormattingDecorator aplica formatação ao relatório.
type FormattingDecorator struct {
	ReportDecorator
}

func (f *FormattingDecorator) Generate() string {
	return fmt.Sprintf("%s with formatting", f.Report.Generate())
}

func (f *FormattingDecorator) GetDescription() string {
	return fmt.Sprintf("%s with Formatting", f.Report.GetDescription())
}
