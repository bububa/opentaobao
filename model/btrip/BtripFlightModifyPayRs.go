package btrip

import (
	"sync"
)

// BtripFlightModifyPayRs 结构体
type BtripFlightModifyPayRs struct {
	// 支付状态
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 支付宝交易流水号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 支付价格
	PayPrice int64 `json:"pay_price,omitempty" xml:"pay_price,omitempty"`
	// 能否重试
	CanRetry bool `json:"can_retry,omitempty" xml:"can_retry,omitempty"`
}

var poolBtripFlightModifyPayRs = sync.Pool{
	New: func() any {
		return new(BtripFlightModifyPayRs)
	},
}

// GetBtripFlightModifyPayRs() 从对象池中获取BtripFlightModifyPayRs
func GetBtripFlightModifyPayRs() *BtripFlightModifyPayRs {
	return poolBtripFlightModifyPayRs.Get().(*BtripFlightModifyPayRs)
}

// ReleaseBtripFlightModifyPayRs 释放BtripFlightModifyPayRs
func ReleaseBtripFlightModifyPayRs(v *BtripFlightModifyPayRs) {
	v.PayStatus = ""
	v.PayTime = ""
	v.TradeNo = ""
	v.PayPrice = 0
	v.CanRetry = false
	poolBtripFlightModifyPayRs.Put(v)
}
