package einvoice

import (
	"sync"
)

// SimpleUnitOrder 结构体
type SimpleUnitOrder struct {
	// 订购单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 税盘编号
	TaxDiskNo string `json:"tax_disk_no,omitempty" xml:"tax_disk_no,omitempty"`
	// 状态-0:待部署,1:部署,2:变更,3:释放
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolSimpleUnitOrder = sync.Pool{
	New: func() any {
		return new(SimpleUnitOrder)
	},
}

// GetSimpleUnitOrder() 从对象池中获取SimpleUnitOrder
func GetSimpleUnitOrder() *SimpleUnitOrder {
	return poolSimpleUnitOrder.Get().(*SimpleUnitOrder)
}

// ReleaseSimpleUnitOrder 释放SimpleUnitOrder
func ReleaseSimpleUnitOrder(v *SimpleUnitOrder) {
	v.OrderId = ""
	v.TaxDiskNo = ""
	v.Status = 0
	poolSimpleUnitOrder.Put(v)
}
