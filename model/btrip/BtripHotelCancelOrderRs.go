package btrip

import (
	"sync"
)

// BtripHotelCancelOrderRs 结构体
type BtripHotelCancelOrderRs struct {
	// 罚金
	ForfeitFee int64 `json:"forfeit_fee,omitempty" xml:"forfeit_fee,omitempty"`
	// 是否取消成功
	CancelSuccess bool `json:"cancel_success,omitempty" xml:"cancel_success,omitempty"`
}

var poolBtripHotelCancelOrderRs = sync.Pool{
	New: func() any {
		return new(BtripHotelCancelOrderRs)
	},
}

// GetBtripHotelCancelOrderRs() 从对象池中获取BtripHotelCancelOrderRs
func GetBtripHotelCancelOrderRs() *BtripHotelCancelOrderRs {
	return poolBtripHotelCancelOrderRs.Get().(*BtripHotelCancelOrderRs)
}

// ReleaseBtripHotelCancelOrderRs 释放BtripHotelCancelOrderRs
func ReleaseBtripHotelCancelOrderRs(v *BtripHotelCancelOrderRs) {
	v.ForfeitFee = 0
	v.CancelSuccess = false
	poolBtripHotelCancelOrderRs.Put(v)
}
