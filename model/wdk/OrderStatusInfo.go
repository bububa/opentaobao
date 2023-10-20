package wdk

import (
	"sync"
)

// OrderStatusInfo 结构体
type OrderStatusInfo struct {
	// 外部渠道店ID(与shop_id必选其一)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 订单状态   已支付： PAID；  TRADE_CLOSE（仅未支付订单）
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 渠道店id(与out_shop_id必选其一)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 经营店Id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道来源(选填out_shop_id时该值必填)
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}

var poolOrderStatusInfo = sync.Pool{
	New: func() any {
		return new(OrderStatusInfo)
	},
}

// GetOrderStatusInfo() 从对象池中获取OrderStatusInfo
func GetOrderStatusInfo() *OrderStatusInfo {
	return poolOrderStatusInfo.Get().(*OrderStatusInfo)
}

// ReleaseOrderStatusInfo 释放OrderStatusInfo
func ReleaseOrderStatusInfo(v *OrderStatusInfo) {
	v.OutShopId = ""
	v.OutOrderId = ""
	v.OrderStatus = ""
	v.ShopId = ""
	v.StoreId = ""
	v.OrderFrom = 0
	poolOrderStatusInfo.Put(v)
}
