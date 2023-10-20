package idle

import (
	"sync"
)

// OrderCreateRequest 结构体
type OrderCreateRequest struct {
	// 额外信息
	ExtraData string `json:"extra_data,omitempty" xml:"extra_data,omitempty"`
	// 虚拟商品code，virtual_item_order为true时必传
	VirtualItemCode string `json:"virtual_item_code,omitempty" xml:"virtual_item_code,omitempty"`
	// 业务产品标识
	XGlobalBizCode string `json:"x_global_biz_code,omitempty" xml:"x_global_biz_code,omitempty"`
	// 邮寄地址的ID。不传入表示不需要邮寄
	DeliverId string `json:"deliver_id,omitempty" xml:"deliver_id,omitempty"`
	// 商品id，virtual_item_order为false时必传
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品单价，单位分。
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 购买数量，默认1个
	BuyQuantity int64 `json:"buy_quantity,omitempty" xml:"buy_quantity,omitempty"`
	// 是否为虚拟商品下单，默认为false
	VirtualItemOrder bool `json:"virtual_item_order,omitempty" xml:"virtual_item_order,omitempty"`
}

var poolOrderCreateRequest = sync.Pool{
	New: func() any {
		return new(OrderCreateRequest)
	},
}

// GetOrderCreateRequest() 从对象池中获取OrderCreateRequest
func GetOrderCreateRequest() *OrderCreateRequest {
	return poolOrderCreateRequest.Get().(*OrderCreateRequest)
}

// ReleaseOrderCreateRequest 释放OrderCreateRequest
func ReleaseOrderCreateRequest(v *OrderCreateRequest) {
	v.ExtraData = ""
	v.VirtualItemCode = ""
	v.XGlobalBizCode = ""
	v.DeliverId = ""
	v.ItemId = 0
	v.Amount = 0
	v.BuyQuantity = 0
	v.VirtualItemOrder = false
	poolOrderCreateRequest.Put(v)
}
