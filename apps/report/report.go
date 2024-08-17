package reportapp

import (
	"diegomorg/go-decorator/apps/report/decorators"
	"diegomorg/go-decorator/apps/report/interfaces"
	"fmt"
)

func logReport(report interfaces.Report) {
	fmt.Printf("%s \n%s \n\n", report.GetDescription(), report.Generate())
}

func Run() {
	// Cria um relatório básico
	var report interfaces.Report = &BaseReport{}
	logReport(report)

	// Adiciona filtros ao relatório básico
	report = &decorators.FilterDecorator{ReportDecorator: decorators.ReportDecorator{Report: report}}
	logReport(report)

	// Adiciona formatação ao relatório com filtros
	report = &decorators.FormattingDecorator{ReportDecorator: decorators.ReportDecorator{Report: report}}
	logReport(report)

	// Adiciona gráficos ao relatório com filtros e formatação
	report = &decorators.GraphDecorator{ReportDecorator: decorators.ReportDecorator{Report: report}}
	logReport(report)

	// Adiciona opções de exportação ao relatório com filtros, formatação e gráficos
	report = &decorators.ExportDecorator{ReportDecorator: decorators.ReportDecorator{Report: report}}
	logReport(report)
}
