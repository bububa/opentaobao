package einvoice

import (
	"sync"
)

// TaxOptimizationSalaryDetailInfoDto 结构体
type TaxOptimizationSalaryDetailInfoDto struct {
	// 承包商编码
	ContractorCode string `json:"contractor_code,omitempty" xml:"contractor_code,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 明细id
	DetailId string `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
	// 用户在业务平台的userid
	IdentificationInBelongingEmployer string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
	// 明细金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolTaxOptimizationSalaryDetailInfoDto = sync.Pool{
	New: func() any {
		return new(TaxOptimizationSalaryDetailInfoDto)
	},
}

// GetTaxOptimizationSalaryDetailInfoDto() 从对象池中获取TaxOptimizationSalaryDetailInfoDto
func GetTaxOptimizationSalaryDetailInfoDto() *TaxOptimizationSalaryDetailInfoDto {
	return poolTaxOptimizationSalaryDetailInfoDto.Get().(*TaxOptimizationSalaryDetailInfoDto)
}

// ReleaseTaxOptimizationSalaryDetailInfoDto 释放TaxOptimizationSalaryDetailInfoDto
func ReleaseTaxOptimizationSalaryDetailInfoDto(v *TaxOptimizationSalaryDetailInfoDto) {
	v.ContractorCode = ""
	v.CreateTime = ""
	v.DetailId = ""
	v.IdentificationInBelongingEmployer = ""
	v.Amount = 0
	poolTaxOptimizationSalaryDetailInfoDto.Put(v)
}
