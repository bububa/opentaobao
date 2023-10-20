package logistic

import (
	"sync"
)

// InvoiceDto 结构体
type InvoiceDto struct {
	// 3-digit number
	InvoiceSeries string `json:"invoice_series,omitempty" xml:"invoice_series,omitempty"`
	// 9-digit number
	InvoiceNumber string `json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// date of inovice issued
	InvoiceDate string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	// value of invoice
	InvoiceTotalValue string `json:"invoice_total_value,omitempty" xml:"invoice_total_value,omitempty"`
	// 44-digit number
	InvoiceKey string `json:"invoice_key,omitempty" xml:"invoice_key,omitempty"`
}

var poolInvoiceDto = sync.Pool{
	New: func() any {
		return new(InvoiceDto)
	},
}

// GetInvoiceDto() 从对象池中获取InvoiceDto
func GetInvoiceDto() *InvoiceDto {
	return poolInvoiceDto.Get().(*InvoiceDto)
}

// ReleaseInvoiceDto 释放InvoiceDto
func ReleaseInvoiceDto(v *InvoiceDto) {
	v.InvoiceSeries = ""
	v.InvoiceNumber = ""
	v.InvoiceDate = ""
	v.InvoiceTotalValue = ""
	v.InvoiceKey = ""
	poolInvoiceDto.Put(v)
}
