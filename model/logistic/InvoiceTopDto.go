package logistic

import (
	"sync"
)

// InvoiceTopDto 结构体
type InvoiceTopDto struct {
	// 3-digit number
	InvoiceSeries string `json:"invoice_series,omitempty" xml:"invoice_series,omitempty"`
	// 44-digit number
	InvoiceKey string `json:"invoice_key,omitempty" xml:"invoice_key,omitempty"`
	// 9-digit number
	InvoiceNumber string `json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// value of invoice
	InvoiceTotalValue string `json:"invoice_total_value,omitempty" xml:"invoice_total_value,omitempty"`
	// date of inovice issued
	InvoiceDate string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
}

var poolInvoiceTopDto = sync.Pool{
	New: func() any {
		return new(InvoiceTopDto)
	},
}

// GetInvoiceTopDto() 从对象池中获取InvoiceTopDto
func GetInvoiceTopDto() *InvoiceTopDto {
	return poolInvoiceTopDto.Get().(*InvoiceTopDto)
}

// ReleaseInvoiceTopDto 释放InvoiceTopDto
func ReleaseInvoiceTopDto(v *InvoiceTopDto) {
	v.InvoiceSeries = ""
	v.InvoiceKey = ""
	v.InvoiceNumber = ""
	v.InvoiceTotalValue = ""
	v.InvoiceDate = ""
	poolInvoiceTopDto.Put(v)
}
