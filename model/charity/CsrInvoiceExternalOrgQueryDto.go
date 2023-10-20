package charity

import (
	"sync"
)

// CsrInvoiceExternalOrgQueryDto 结构体
type CsrInvoiceExternalOrgQueryDto struct {
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolCsrInvoiceExternalOrgQueryDto = sync.Pool{
	New: func() any {
		return new(CsrInvoiceExternalOrgQueryDto)
	},
}

// GetCsrInvoiceExternalOrgQueryDto() 从对象池中获取CsrInvoiceExternalOrgQueryDto
func GetCsrInvoiceExternalOrgQueryDto() *CsrInvoiceExternalOrgQueryDto {
	return poolCsrInvoiceExternalOrgQueryDto.Get().(*CsrInvoiceExternalOrgQueryDto)
}

// ReleaseCsrInvoiceExternalOrgQueryDto 释放CsrInvoiceExternalOrgQueryDto
func ReleaseCsrInvoiceExternalOrgQueryDto(v *CsrInvoiceExternalOrgQueryDto) {
	v.PageIndex = 0
	v.PageSize = 0
	poolCsrInvoiceExternalOrgQueryDto.Put(v)
}
