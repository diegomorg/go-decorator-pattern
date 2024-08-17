package decorators

import (
	"diegomorg/go-decorator/apps/report/interfaces"
)

// ReportDecorator é uma estrutura base para os decoradores de relatórios.
type ReportDecorator struct {
	Report interfaces.Report
}

func (d *ReportDecorator) Generate() string {
	return d.Report.Generate()
}

func (d *ReportDecorator) GetDescription() string {
	return d.Report.GetDescription()
}
