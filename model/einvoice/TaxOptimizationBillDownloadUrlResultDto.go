package einvoice

import (
	"sync"
)

// TaxOptimizationBillDownloadUrlResultDto 结构体
type TaxOptimizationBillDownloadUrlResultDto struct {
	// 账单文件的下载地址，请求成功后20s内有效
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
}

var poolTaxOptimizationBillDownloadUrlResultDto = sync.Pool{
	New: func() any {
		return new(TaxOptimizationBillDownloadUrlResultDto)
	},
}

// GetTaxOptimizationBillDownloadUrlResultDto() 从对象池中获取TaxOptimizationBillDownloadUrlResultDto
func GetTaxOptimizationBillDownloadUrlResultDto() *TaxOptimizationBillDownloadUrlResultDto {
	return poolTaxOptimizationBillDownloadUrlResultDto.Get().(*TaxOptimizationBillDownloadUrlResultDto)
}

// ReleaseTaxOptimizationBillDownloadUrlResultDto 释放TaxOptimizationBillDownloadUrlResultDto
func ReleaseTaxOptimizationBillDownloadUrlResultDto(v *TaxOptimizationBillDownloadUrlResultDto) {
	v.DownloadUrl = ""
	poolTaxOptimizationBillDownloadUrlResultDto.Put(v)
}
