package charity

// CsrInvoiceApplyProjectDto 结构体
type CsrInvoiceApplyProjectDto struct {
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 项目金额，单位分
	ProjectAmount int64 `json:"project_amount,omitempty" xml:"project_amount,omitempty"`
}
