package einvoice

import (
	"sync"
)

// OrderRightsInfo 结构体
type OrderRightsInfo struct {
	// 订单号
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 订单应开票时间
	ExceptInvoiceTime string `json:"except_invoice_time,omitempty" xml:"except_invoice_time,omitempty"`
}

var poolOrderRightsInfo = sync.Pool{
	New: func() any {
		return new(OrderRightsInfo)
	},
}

// GetOrderRightsInfo() 从对象池中获取OrderRightsInfo
func GetOrderRightsInfo() *OrderRightsInfo {
	return poolOrderRightsInfo.Get().(*OrderRightsInfo)
}

// ReleaseOrderRightsInfo 释放OrderRightsInfo
func ReleaseOrderRightsInfo(v *OrderRightsInfo) {
	v.Tid = ""
	v.ExceptInvoiceTime = ""
	poolOrderRightsInfo.Put(v)
}
