package wdk

import (
	"sync"
)

// InboundItemInfo 结构体
type InboundItemInfo struct {
	// 收货数量
	InboundQuantity string `json:"inbound_quantity,omitempty" xml:"inbound_quantity,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}

var poolInboundItemInfo = sync.Pool{
	New: func() any {
		return new(InboundItemInfo)
	},
}

// GetInboundItemInfo() 从对象池中获取InboundItemInfo
func GetInboundItemInfo() *InboundItemInfo {
	return poolInboundItemInfo.Get().(*InboundItemInfo)
}

// ReleaseInboundItemInfo 释放InboundItemInfo
func ReleaseInboundItemInfo(v *InboundItemInfo) {
	v.InboundQuantity = ""
	v.SkuCode = ""
	poolInboundItemInfo.Put(v)
}
