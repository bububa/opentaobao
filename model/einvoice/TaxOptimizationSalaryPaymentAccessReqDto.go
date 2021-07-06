package einvoice

// TaxOptimizationSalaryPaymentAccessReqDto 结构体
type TaxOptimizationSalaryPaymentAccessReqDto struct {
	// 发薪明细列表
	DetailIdList []string `json:"detail_id_list,omitempty" xml:"detail_id_list>string,omitempty"`
	// 账期
	AccountDate string `json:"account_date,omitempty" xml:"account_date,omitempty"`
	// 发薪金额
	ApplyAmount string `json:"apply_amount,omitempty" xml:"apply_amount,omitempty"`
	// 业务时间
	BusinessTime string `json:"business_time,omitempty" xml:"business_time,omitempty"`
	// 业务方编码
	EmployerCode string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
	// 用户在业务方平台的userid
	IdentificationInBelongingEmployer string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
	// 发薪模式
	PaySalaryMode string `json:"pay_salary_mode,omitempty" xml:"pay_salary_mode,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
