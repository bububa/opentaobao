package btrip

import (
	"sync"
)

// InvoiceInfo 结构体
type InvoiceInfo struct {
	// 发票抬头
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 发票id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
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
	v.Title = ""
	v.Id = 0
	poolInvoiceInfo.Put(v)
}
