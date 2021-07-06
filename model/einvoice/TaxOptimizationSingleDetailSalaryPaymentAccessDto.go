package einvoice

// TaxOptimizationSingleDetailSalaryPaymentAccessDto 结构体
type TaxOptimizationSingleDetailSalaryPaymentAccessDto struct {
	// 业务提交时间
	BusinessTime string `json:"business_time,omitempty" xml:"business_time,omitempty"`
	// 承包商编码
	ContractorCode string `json:"contractor_code,omitempty" xml:"contractor_code,omitempty"`
	// 发薪流水id
	DetailId string `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
	// 业务方编码
	EmployerCode string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
	// 用户在业务方的userid
	IdentificationInBelongingEmployer string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
	// 业务自定义发薪备注
	SalaryRemark string `json:"salary_remark,omitempty" xml:"salary_remark,omitempty"`
	// 发薪金额
	ApplyAmount int64 `json:"apply_amount,omitempty" xml:"apply_amount,omitempty"`
}
