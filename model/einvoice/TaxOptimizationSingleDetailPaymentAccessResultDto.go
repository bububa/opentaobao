package einvoice

import (
	"sync"
)

// TaxOptimizationSingleDetailPaymentAccessResultDto 结构体
type TaxOptimizationSingleDetailPaymentAccessResultDto struct {
	// 服务返回结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaxOptimizationSingleDetailPaymentAccessResultDto = sync.Pool{
	New: func() any {
		return new(TaxOptimizationSingleDetailPaymentAccessResultDto)
	},
}

// GetTaxOptimizationSingleDetailPaymentAccessResultDto() 从对象池中获取TaxOptimizationSingleDetailPaymentAccessResultDto
func GetTaxOptimizationSingleDetailPaymentAccessResultDto() *TaxOptimizationSingleDetailPaymentAccessResultDto {
	return poolTaxOptimizationSingleDetailPaymentAccessResultDto.Get().(*TaxOptimizationSingleDetailPaymentAccessResultDto)
}

// ReleaseTaxOptimizationSingleDetailPaymentAccessResultDto 释放TaxOptimizationSingleDetailPaymentAccessResultDto
func ReleaseTaxOptimizationSingleDetailPaymentAccessResultDto(v *TaxOptimizationSingleDetailPaymentAccessResultDto) {
	v.Success = false
	poolTaxOptimizationSingleDetailPaymentAccessResultDto.Put(v)
}
