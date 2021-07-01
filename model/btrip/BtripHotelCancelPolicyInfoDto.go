package btrip

// BtripHotelCancelPolicyInfoDto 结构体
type BtripHotelCancelPolicyInfoDto struct {
	// 提前小时
	Hour int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// 罚金
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}
