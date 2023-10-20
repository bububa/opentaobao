package charity

import (
	"sync"
)

// CsrInvoiceExternalOrgDrawDto 结构体
type CsrInvoiceExternalOrgDrawDto struct {
	// 发票文件
	FileList []CsrInvoiceFileDto `json:"file_list,omitempty" xml:"file_list>csr_invoice_file_dto,omitempty"`
	// 发票id
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 开票操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
}

var poolCsrInvoiceExternalOrgDrawDto = sync.Pool{
	New: func() any {
		return new(CsrInvoiceExternalOrgDrawDto)
	},
}

// GetCsrInvoiceExternalOrgDrawDto() 从对象池中获取CsrInvoiceExternalOrgDrawDto
func GetCsrInvoiceExternalOrgDrawDto() *CsrInvoiceExternalOrgDrawDto {
	return poolCsrInvoiceExternalOrgDrawDto.Get().(*CsrInvoiceExternalOrgDrawDto)
}

// ReleaseCsrInvoiceExternalOrgDrawDto 释放CsrInvoiceExternalOrgDrawDto
func ReleaseCsrInvoiceExternalOrgDrawDto(v *CsrInvoiceExternalOrgDrawDto) {
	v.FileList = v.FileList[:0]
	v.InvoiceId = ""
	v.Operator = ""
	poolCsrInvoiceExternalOrgDrawDto.Put(v)
}
