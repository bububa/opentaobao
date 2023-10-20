package btrip

import (
	"sync"
)

// OpenInvoiceDo 结构体
type OpenInvoiceDo struct {
	// 发票抬头
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 第三方发票id
	ThirdPartInvoiceId string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// 商旅发票id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 发票类型：1:增值税普通发票；2:增值税专用发票
	InvoiceType int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
}

var poolOpenInvoiceDo = sync.Pool{
	New: func() any {
		return new(OpenInvoiceDo)
	},
}

// GetOpenInvoiceDo() 从对象池中获取OpenInvoiceDo
func GetOpenInvoiceDo() *OpenInvoiceDo {
	return poolOpenInvoiceDo.Get().(*OpenInvoiceDo)
}

// ReleaseOpenInvoiceDo 释放OpenInvoiceDo
func ReleaseOpenInvoiceDo(v *OpenInvoiceDo) {
	v.Title = ""
	v.ThirdPartInvoiceId = ""
	v.Id = 0
	v.InvoiceType = 0
	poolOpenInvoiceDo.Put(v)
}
