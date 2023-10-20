package charity

import (
	"sync"
)

// CsrInvoiceApplyProjectDto 结构体
type CsrInvoiceApplyProjectDto struct {
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 项目金额，单位分
	ProjectAmount int64 `json:"project_amount,omitempty" xml:"project_amount,omitempty"`
}

var poolCsrInvoiceApplyProjectDto = sync.Pool{
	New: func() any {
		return new(CsrInvoiceApplyProjectDto)
	},
}

// GetCsrInvoiceApplyProjectDto() 从对象池中获取CsrInvoiceApplyProjectDto
func GetCsrInvoiceApplyProjectDto() *CsrInvoiceApplyProjectDto {
	return poolCsrInvoiceApplyProjectDto.Get().(*CsrInvoiceApplyProjectDto)
}

// ReleaseCsrInvoiceApplyProjectDto 释放CsrInvoiceApplyProjectDto
func ReleaseCsrInvoiceApplyProjectDto(v *CsrInvoiceApplyProjectDto) {
	v.ProjectName = ""
	v.ProjectAmount = 0
	poolCsrInvoiceApplyProjectDto.Put(v)
}
