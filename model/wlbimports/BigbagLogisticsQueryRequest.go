package wlbimports

// BigbagLogisticsQueryRequest 结构体
type BigbagLogisticsQueryRequest struct {
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 大包id
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
}
