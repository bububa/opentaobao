package wdk

import (
	"sync"
)

// OutboundItemInfo 结构体
type OutboundItemInfo struct {
	// 容器信息
	Containers []ContainerDo `json:"containers,omitempty" xml:"containers>container_do,omitempty"`
	// 批发单号
	WholesaleOrderNo string `json:"wholesale_order_no,omitempty" xml:"wholesale_order_no,omitempty"`
	// 已废弃，请使用containers.production_date
	ProductionDate string `json:"production_date,omitempty" xml:"production_date,omitempty"`
	// 出库数量(为正数或零)
	OutboundQuantity string `json:"outbound_quantity,omitempty" xml:"outbound_quantity,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 外部单号，如采购单号
	ExternalOrderNo string `json:"external_order_no,omitempty" xml:"external_order_no,omitempty"`
	// 是否完结
	OutboundCompleted bool `json:"outbound_completed,omitempty" xml:"outbound_completed,omitempty"`
}

var poolOutboundItemInfo = sync.Pool{
	New: func() any {
		return new(OutboundItemInfo)
	},
}

// GetOutboundItemInfo() 从对象池中获取OutboundItemInfo
func GetOutboundItemInfo() *OutboundItemInfo {
	return poolOutboundItemInfo.Get().(*OutboundItemInfo)
}

// ReleaseOutboundItemInfo 释放OutboundItemInfo
func ReleaseOutboundItemInfo(v *OutboundItemInfo) {
	v.Containers = v.Containers[:0]
	v.WholesaleOrderNo = ""
	v.ProductionDate = ""
	v.OutboundQuantity = ""
	v.SkuCode = ""
	v.ExternalOrderNo = ""
	v.OutboundCompleted = false
	poolOutboundItemInfo.Put(v)
}
