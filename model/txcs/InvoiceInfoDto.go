package txcs

import (
	"sync"
)

// InvoiceInfoDto 结构体
type InvoiceInfoDto struct {
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
}

var poolInvoiceInfoDto = sync.Pool{
	New: func() any {
		return new(InvoiceInfoDto)
	},
}

// GetInvoiceInfoDto() 从对象池中获取InvoiceInfoDto
func GetInvoiceInfoDto() *InvoiceInfoDto {
	return poolInvoiceInfoDto.Get().(*InvoiceInfoDto)
}

// ReleaseInvoiceInfoDto 释放InvoiceInfoDto
func ReleaseInvoiceInfoDto(v *InvoiceInfoDto) {
	v.InvoiceNo = ""
	v.InvoiceCode = ""
	poolInvoiceInfoDto.Put(v)
}
