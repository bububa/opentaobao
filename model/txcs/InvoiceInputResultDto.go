package txcs

import (
	"sync"
)

// InvoiceInputResultDto 结构体
type InvoiceInputResultDto struct {
	// 失败原因
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
}

var poolInvoiceInputResultDto = sync.Pool{
	New: func() any {
		return new(InvoiceInputResultDto)
	},
}

// GetInvoiceInputResultDto() 从对象池中获取InvoiceInputResultDto
func GetInvoiceInputResultDto() *InvoiceInputResultDto {
	return poolInvoiceInputResultDto.Get().(*InvoiceInputResultDto)
}

// ReleaseInvoiceInputResultDto 释放InvoiceInputResultDto
func ReleaseInvoiceInputResultDto(v *InvoiceInputResultDto) {
	v.Msg = ""
	v.InvoiceCode = ""
	v.InvoiceNo = ""
	poolInvoiceInputResultDto.Put(v)
}
