package pur

// AccessLadderPriceDto 结构体
type AccessLadderPriceDto struct {
	// 原价
	OriginUnitPrice float64 `json:"origin_unit_price,omitempty" xml:"origin_unit_price,omitempty"`
	// 协议价
	UnitPrice float64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 最小采购量
	MinimumPurchaseQuantity float64 `json:"minimum_purchase_quantity,omitempty" xml:"minimum_purchase_quantity,omitempty"`
}
