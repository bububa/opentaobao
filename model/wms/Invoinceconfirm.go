package wms

import (
	"sync"
)

// Invoinceconfirm 结构体
type Invoinceconfirm struct {
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 发票号
	InvoiceNumber string `json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// Erp发票ID
	BillId string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
}

var poolInvoinceconfirm = sync.Pool{
	New: func() any {
		return new(Invoinceconfirm)
	},
}

// GetInvoinceconfirm() 从对象池中获取Invoinceconfirm
func GetInvoinceconfirm() *Invoinceconfirm {
	return poolInvoinceconfirm.Get().(*Invoinceconfirm)
}

// ReleaseInvoinceconfirm 释放Invoinceconfirm
func ReleaseInvoinceconfirm(v *Invoinceconfirm) {
	v.InvoiceCode = ""
	v.InvoiceNumber = ""
	v.BillId = ""
	poolInvoinceconfirm.Put(v)
}
