package einvoice

import (
	"sync"
)

// TaxOptimizationQueryPaySalaryAccountDto 结构体
type TaxOptimizationQueryPaySalaryAccountDto struct {
	// 承包商编码
	ContractorCode string `json:"contractor_code,omitempty" xml:"contractor_code,omitempty"`
	// 业务方编码
	EmployerCode string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
	// 用户在业务方平台的userid
	IdentificationInBelongingEmployer string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
	// 税优模式
	TaxOptimizationMode string `json:"tax_optimization_mode,omitempty" xml:"tax_optimization_mode,omitempty"`
}

var poolTaxOptimizationQueryPaySalaryAccountDto = sync.Pool{
	New: func() any {
		return new(TaxOptimizationQueryPaySalaryAccountDto)
	},
}

// GetTaxOptimizationQueryPaySalaryAccountDto() 从对象池中获取TaxOptimizationQueryPaySalaryAccountDto
func GetTaxOptimizationQueryPaySalaryAccountDto() *TaxOptimizationQueryPaySalaryAccountDto {
	return poolTaxOptimizationQueryPaySalaryAccountDto.Get().(*TaxOptimizationQueryPaySalaryAccountDto)
}

// ReleaseTaxOptimizationQueryPaySalaryAccountDto 释放TaxOptimizationQueryPaySalaryAccountDto
func ReleaseTaxOptimizationQueryPaySalaryAccountDto(v *TaxOptimizationQueryPaySalaryAccountDto) {
	v.ContractorCode = ""
	v.EmployerCode = ""
	v.IdentificationInBelongingEmployer = ""
	v.TaxOptimizationMode = ""
	poolTaxOptimizationQueryPaySalaryAccountDto.Put(v)
}
