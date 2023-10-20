package einvoice

import (
	"sync"
)

// TaxOptimizationSalaryPaymentAccessResultDto 结构体
type TaxOptimizationSalaryPaymentAccessResultDto struct {
	// 发薪状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 失败的个数
	FailCount int64 `json:"fail_count,omitempty" xml:"fail_count,omitempty"`
	// 正在处理的个数
	ProcessingCount int64 `json:"processing_count,omitempty" xml:"processing_count,omitempty"`
	// 成功的个数
	SuccessCount int64 `json:"success_count,omitempty" xml:"success_count,omitempty"`
	// 总的发薪个数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolTaxOptimizationSalaryPaymentAccessResultDto = sync.Pool{
	New: func() any {
		return new(TaxOptimizationSalaryPaymentAccessResultDto)
	},
}

// GetTaxOptimizationSalaryPaymentAccessResultDto() 从对象池中获取TaxOptimizationSalaryPaymentAccessResultDto
func GetTaxOptimizationSalaryPaymentAccessResultDto() *TaxOptimizationSalaryPaymentAccessResultDto {
	return poolTaxOptimizationSalaryPaymentAccessResultDto.Get().(*TaxOptimizationSalaryPaymentAccessResultDto)
}

// ReleaseTaxOptimizationSalaryPaymentAccessResultDto 释放TaxOptimizationSalaryPaymentAccessResultDto
func ReleaseTaxOptimizationSalaryPaymentAccessResultDto(v *TaxOptimizationSalaryPaymentAccessResultDto) {
	v.Status = ""
	v.FailCount = 0
	v.ProcessingCount = 0
	v.SuccessCount = 0
	v.TotalCount = 0
	poolTaxOptimizationSalaryPaymentAccessResultDto.Put(v)
}
