package reportapp

// BaseReport é uma implementação concreta da interface Report.
type BaseReport struct{}

func (r *BaseReport) Generate() string {
	return "Report content"
}

func (r *BaseReport) GetDescription() string {
	return "Base Report"
}
