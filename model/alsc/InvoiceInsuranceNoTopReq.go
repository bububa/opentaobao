package alsc

import (
	"sync"
)

// InvoiceInsuranceNoTopReq 结构体
type InvoiceInsuranceNoTopReq struct {
	// 口碑开票id
	InvoiceApplyId string `json:"invoice_apply_id,omitempty" xml:"invoice_apply_id,omitempty"`
}

var poolInvoiceInsuranceNoTopReq = sync.Pool{
	New: func() any {
		return new(InvoiceInsuranceNoTopReq)
	},
}

// GetInvoiceInsuranceNoTopReq() 从对象池中获取InvoiceInsuranceNoTopReq
func GetInvoiceInsuranceNoTopReq() *InvoiceInsuranceNoTopReq {
	return poolInvoiceInsuranceNoTopReq.Get().(*InvoiceInsuranceNoTopReq)
}

// ReleaseInvoiceInsuranceNoTopReq 释放InvoiceInsuranceNoTopReq
func ReleaseInvoiceInsuranceNoTopReq(v *InvoiceInsuranceNoTopReq) {
	v.InvoiceApplyId = ""
	poolInvoiceInsuranceNoTopReq.Put(v)
}
