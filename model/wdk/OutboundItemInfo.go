package wdk

// OutboundItemInfo 结构体
type OutboundItemInfo struct {
	// 批发单号
	WholesaleOrderNo string `json:"wholesale_order_no,omitempty" xml:"wholesale_order_no,omitempty"`
	// 已废弃，请使用containers.production_date
	ProductionDate string `json:"production_date,omitempty" xml:"production_date,omitempty"`
	// 是否完结
	OutboundCompleted bool `json:"outbound_completed,omitempty" xml:"outbound_completed,omitempty"`
	// 出库数量(为正数或零)
	OutboundQuantity string `json:"outbound_quantity,omitempty" xml:"outbound_quantity,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 外部单号，如采购单号
	ExternalOrderNo string `json:"external_order_no,omitempty" xml:"external_order_no,omitempty"`
	// 容器信息
	Containers []ContainerDo `json:"containers,omitempty" xml:"containers>container_do,omitempty"`
}
