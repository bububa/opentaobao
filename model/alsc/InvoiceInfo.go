package alsc

import (
	"sync"
)

// InvoiceInfo 结构体
type InvoiceInfo struct {
	// 发票抬头
	Invoice string `json:"invoice,omitempty" xml:"invoice,omitempty"`
	// 发票类型：  个人-PERSONAL公司-COMPANY
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 纳税人识别码
	TaxPayerId string `json:"tax_payer_id,omitempty" xml:"tax_payer_id,omitempty"`
}

var poolInvoiceInfo = sync.Pool{
	New: func() any {
		return new(InvoiceInfo)
	},
}

// GetInvoiceInfo() 从对象池中获取InvoiceInfo
func GetInvoiceInfo() *InvoiceInfo {
	return poolInvoiceInfo.Get().(*InvoiceInfo)
}

// ReleaseInvoiceInfo 释放InvoiceInfo
func ReleaseInvoiceInfo(v *InvoiceInfo) {
	v.Invoice = ""
	v.InvoiceType = ""
	v.TaxPayerId = ""
	poolInvoiceInfo.Put(v)
}
