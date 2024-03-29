package ascpchannel

import (
	"sync"
)

// Orderitems 结构体
type Orderitems struct {
	// 销退回告明细列表
	InstorageDetails []Instoragedetails `json:"instorage_details,omitempty" xml:"instorage_details>instoragedetails,omitempty"`
	// 履约子单号
	SubOrderCode string `json:"sub_order_code,omitempty" xml:"sub_order_code,omitempty"`
	// 货品ID
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// erp订单明细行号
	ErpOrderLine string `json:"erp_order_line,omitempty" xml:"erp_order_line,omitempty"`
	// 货品实发数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 货品缺发数量
	LackQuantity int64 `json:"lack_quantity,omitempty" xml:"lack_quantity,omitempty"`
	// 货品计划退回数量
	PlanReturnQuantity int64 `json:"plan_return_quantity,omitempty" xml:"plan_return_quantity,omitempty"`
	// 货品实际收货总数量
	ActualReceivedQuantity int64 `json:"actual_received_quantity,omitempty" xml:"actual_received_quantity,omitempty"`
	// 货品未收货总数量
	ActualLackQuantity int64 `json:"actual_lack_quantity,omitempty" xml:"actual_lack_quantity,omitempty"`
}

var poolOrderitems = sync.Pool{
	New: func() any {
		return new(Orderitems)
	},
}

// GetOrderitems() 从对象池中获取Orderitems
func GetOrderitems() *Orderitems {
	return poolOrderitems.Get().(*Orderitems)
}

// ReleaseOrderitems 释放Orderitems
func ReleaseOrderitems(v *Orderitems) {
	v.InstorageDetails = v.InstorageDetails[:0]
	v.SubOrderCode = ""
	v.ScItemId = ""
	v.ErpOrderLine = ""
	v.ItemQuantity = 0
	v.LackQuantity = 0
	v.PlanReturnQuantity = 0
	v.ActualReceivedQuantity = 0
	v.ActualLackQuantity = 0
	poolOrderitems.Put(v)
}
