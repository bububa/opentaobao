package einvoice

// AgreementInfoDto 结构体
type AgreementInfoDto struct {
	// 协议类型
	AgreementType string `json:"agreement_type,omitempty" xml:"agreement_type,omitempty"`
	// 资产符号
	AssetSymbol string `json:"asset_symbol,omitempty" xml:"asset_symbol,omitempty"`
	// 业务方编码
	EmployerCode string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
	// 签约时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 合同链接
	AgreementUrl string `json:"agreement_url,omitempty" xml:"agreement_url,omitempty"`
	// 供应商编码
	ContractorCode string `json:"contractor_code,omitempty" xml:"contractor_code,omitempty"`
	// 供应商名称
	ContractorName string `json:"contractor_name,omitempty" xml:"contractor_name,omitempty"`
	// 资产类型
	AssetType string `json:"asset_type,omitempty" xml:"asset_type,omitempty"`
	// 报税类型
	ApplyDutiableModeEnum string `json:"apply_dutiable_mode_enum,omitempty" xml:"apply_dutiable_mode_enum,omitempty"`
	// 发薪模式
	PaySalaryModeEnum string `json:"pay_salary_mode_enum,omitempty" xml:"pay_salary_mode_enum,omitempty"`
	// 税优模式
	TaxOptimizationMode string `json:"tax_optimization_mode,omitempty" xml:"tax_optimization_mode,omitempty"`
	// 解约时间
	TerminationTime string `json:"termination_time,omitempty" xml:"termination_time,omitempty"`
	// 用户在业务方平台的userid
	IdentificationInBelongingEmployer string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
	// 扩展信息
	ExtendField string `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	// 签约状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
