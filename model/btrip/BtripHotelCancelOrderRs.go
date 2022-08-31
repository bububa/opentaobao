package btrip

// BtripHotelCancelOrderRs 结构体
type BtripHotelCancelOrderRs struct {
	// 罚金
	ForfeitFee int64 `json:"forfeit_fee,omitempty" xml:"forfeit_fee,omitempty"`
	// 是否取消成功
	CancelSuccess bool `json:"cancel_success,omitempty" xml:"cancel_success,omitempty"`
}
