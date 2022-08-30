package wlbimports

// BigbagStatusRequest 结构体
type BigbagStatusRequest struct {
	// 大包id
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
