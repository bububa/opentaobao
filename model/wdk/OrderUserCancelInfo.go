package wdk

import (
	"sync"
)

// OrderUserCancelInfo 结构体
type OrderUserCancelInfo struct {
	// 外部订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部渠道店ID(与shop_id必选其一)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 渠道店id(与out_shop_id必选其一)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 盒马主单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 经营店Id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道来源(选填out_shop_id时该值必填)
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}

var poolOrderUserCancelInfo = sync.Pool{
	New: func() any {
		return new(OrderUserCancelInfo)
	},
}

// GetOrderUserCancelInfo() 从对象池中获取OrderUserCancelInfo
func GetOrderUserCancelInfo() *OrderUserCancelInfo {
	return poolOrderUserCancelInfo.Get().(*OrderUserCancelInfo)
}

// ReleaseOrderUserCancelInfo 释放OrderUserCancelInfo
func ReleaseOrderUserCancelInfo(v *OrderUserCancelInfo) {
	v.OutOrderId = ""
	v.OutShopId = ""
	v.ShopId = ""
	v.BizOrderId = ""
	v.StoreId = ""
	v.OrderFrom = 0
	poolOrderUserCancelInfo.Put(v)
}
