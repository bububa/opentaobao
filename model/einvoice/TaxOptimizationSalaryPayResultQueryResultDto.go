package einvoice

import (
	"sync"
)

// TaxOptimizationSalaryPayResultQueryResultDto 结构体
type TaxOptimizationSalaryPayResultQueryResultDto struct {
	// 结果列表
	SalaryDetailList []SalaryDetailDto `json:"salary_detail_list,omitempty" xml:"salary_detail_list>salary_detail_dto,omitempty"`
}

var poolTaxOptimizationSalaryPayResultQueryResultDto = sync.Pool{
	New: func() any {
		return new(TaxOptimizationSalaryPayResultQueryResultDto)
	},
}

// GetTaxOptimizationSalaryPayResultQueryResultDto() 从对象池中获取TaxOptimizationSalaryPayResultQueryResultDto
func GetTaxOptimizationSalaryPayResultQueryResultDto() *TaxOptimizationSalaryPayResultQueryResultDto {
	return poolTaxOptimizationSalaryPayResultQueryResultDto.Get().(*TaxOptimizationSalaryPayResultQueryResultDto)
}

// ReleaseTaxOptimizationSalaryPayResultQueryResultDto 释放TaxOptimizationSalaryPayResultQueryResultDto
func ReleaseTaxOptimizationSalaryPayResultQueryResultDto(v *TaxOptimizationSalaryPayResultQueryResultDto) {
	v.SalaryDetailList = v.SalaryDetailList[:0]
	poolTaxOptimizationSalaryPayResultQueryResultDto.Put(v)
}
