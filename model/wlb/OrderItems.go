package wlb

import (
	"sync"
)

// OrderItems 结构体
type OrderItems struct {
	// 明细对应主单的交易单号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 明细对应的子交易单号
	TradeItemId string `json:"trade_item_id,omitempty" xml:"trade_item_id,omitempty"`
	// 默认：0；促销赠品1001
	ItemTag string `json:"item_tag,omitempty" xml:"item_tag,omitempty"`
	// 货品id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 前端商家编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 前端宝贝itemId
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 前端skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 后端商家编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 明细ID
	OrderItemId int64 `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 商品金额 123.33元，单位：分
	ItemAmount int64 `json:"item_amount,omitempty" xml:"item_amount,omitempty"`
}

var poolOrderItems = sync.Pool{
	New: func() any {
		return new(OrderItems)
	},
}

// GetOrderItems() 从对象池中获取OrderItems
func GetOrderItems() *OrderItems {
	return poolOrderItems.Get().(*OrderItems)
}

// ReleaseOrderItems 释放OrderItems
func ReleaseOrderItems(v *OrderItems) {
	v.TradeId = ""
	v.TradeItemId = ""
	v.ItemTag = ""
	v.ScItemId = ""
	v.ItemCode = ""
	v.ItemId = ""
	v.SkuId = ""
	v.ScItemCode = ""
	v.OrderItemId = 0
	v.ItemQuantity = 0
	v.ItemAmount = 0
	poolOrderItems.Put(v)
}
