package einvoice

// SalaryDetailDto 结构体
type SalaryDetailDto struct {
	// 账期
	AccountDate string `json:"account_date,omitempty" xml:"account_date,omitempty"`
	// 收款账号
	AssetSymbol string `json:"asset_symbol,omitempty" xml:"asset_symbol,omitempty"`
	// 账号类型
	AssetType string `json:"asset_type,omitempty" xml:"asset_type,omitempty"`
	// 业务时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 明细流水号
	DetailId string `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
	// 业务编码
	EmployerCode string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
	// 发薪失败原因
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 用户在业务方平台上的userid
	IdentificationInBelongingEmployer string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
	// 执行时间
	ProcessTime string `json:"process_time,omitempty" xml:"process_time,omitempty"`
	// 发薪状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 发薪金额
	ApplyAmount int64 `json:"apply_amount,omitempty" xml:"apply_amount,omitempty"`
	// 已发金额
	SalaryAmount int64 `json:"salary_amount,omitempty" xml:"salary_amount,omitempty"`
}
