package einvoice

import (
	"sync"
)

// InvoiceItems 结构体
type InvoiceItems struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 价税合计
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 规格型号
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
}

var poolInvoiceItems = sync.Pool{
	New: func() any {
		return new(InvoiceItems)
	},
}

// GetInvoiceItems() 从对象池中获取InvoiceItems
func GetInvoiceItems() *InvoiceItems {
	return poolInvoiceItems.Get().(*InvoiceItems)
}

// ReleaseInvoiceItems 释放InvoiceItems
func ReleaseInvoiceItems(v *InvoiceItems) {
	v.ItemName = ""
	v.Quantity = ""
	v.Amount = ""
	v.Specification = ""
	v.Unit = ""
	poolInvoiceItems.Put(v)
}
