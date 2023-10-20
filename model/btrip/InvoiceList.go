package btrip

import (
	"sync"
)

// InvoiceList 结构体
type InvoiceList struct {
	// title
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolInvoiceList = sync.Pool{
	New: func() any {
		return new(InvoiceList)
	},
}

// GetInvoiceList() 从对象池中获取InvoiceList
func GetInvoiceList() *InvoiceList {
	return poolInvoiceList.Get().(*InvoiceList)
}

// ReleaseInvoiceList 释放InvoiceList
func ReleaseInvoiceList(v *InvoiceList) {
	v.Title = ""
	v.Id = 0
	poolInvoiceList.Put(v)
}
