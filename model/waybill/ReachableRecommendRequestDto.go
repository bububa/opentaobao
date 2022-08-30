package waybill

// ReachableRecommendRequestDto 结构体
type ReachableRecommendRequestDto struct {
	// 指定快递公司
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 地址
	Address *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
}
