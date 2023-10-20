package charity

import (
	"sync"
)

// CsrInvoiceExternalOrgRejectDto 结构体
type CsrInvoiceExternalOrgRejectDto struct {
	// 拒绝原因
	RejectReason string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	// 发票id
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
}

var poolCsrInvoiceExternalOrgRejectDto = sync.Pool{
	New: func() any {
		return new(CsrInvoiceExternalOrgRejectDto)
	},
}

// GetCsrInvoiceExternalOrgRejectDto() 从对象池中获取CsrInvoiceExternalOrgRejectDto
func GetCsrInvoiceExternalOrgRejectDto() *CsrInvoiceExternalOrgRejectDto {
	return poolCsrInvoiceExternalOrgRejectDto.Get().(*CsrInvoiceExternalOrgRejectDto)
}

// ReleaseCsrInvoiceExternalOrgRejectDto 释放CsrInvoiceExternalOrgRejectDto
func ReleaseCsrInvoiceExternalOrgRejectDto(v *CsrInvoiceExternalOrgRejectDto) {
	v.RejectReason = ""
	v.InvoiceId = ""
	v.Operator = ""
	poolCsrInvoiceExternalOrgRejectDto.Put(v)
}
