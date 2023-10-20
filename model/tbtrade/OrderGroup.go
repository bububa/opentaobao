package tbtrade

import (
	"sync"
)

// OrderGroup 结构体
type OrderGroup struct {
	// 淘宝交易子订单id
	TaobaoSubOrderId string `json:"taobao_sub_order_id,omitempty" xml:"taobao_sub_order_id,omitempty"`
	// 淘宝单号可以是交易订单、换货单、补货单、代发单或分销单等，当 order_type=0时，order_id =  taobao_parent_order_id。
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// erp系统内的订单id
	ErpOrderId string `json:"erp_order_id,omitempty" xml:"erp_order_id,omitempty"`
	// 淘宝交易主订单id
	TaobaoParentOrderId string `json:"taobao_parent_order_id,omitempty" xml:"taobao_parent_order_id,omitempty"`
	// order_id的类型，0:淘宝交易订单,1: 换货单,2:分销单，3:补货单，4:代发单
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 商品类型, 0:下单货品，1:赠品，2:其他
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
}

var poolOrderGroup = sync.Pool{
	New: func() any {
		return new(OrderGroup)
	},
}

// GetOrderGroup() 从对象池中获取OrderGroup
func GetOrderGroup() *OrderGroup {
	return poolOrderGroup.Get().(*OrderGroup)
}

// ReleaseOrderGroup 释放OrderGroup
func ReleaseOrderGroup(v *OrderGroup) {
	v.TaobaoSubOrderId = ""
	v.OrderId = ""
	v.ErpOrderId = ""
	v.TaobaoParentOrderId = ""
	v.OrderType = 0
	v.ItemType = 0
	poolOrderGroup.Put(v)
}
