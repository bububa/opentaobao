package btrip

import (
	"sync"
)

// BtripFlightPayOrderRs 结构体
type BtripFlightPayOrderRs struct {
	// 支付宝流水号（现付支付宝不为空）
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 最晚支付时间
	LastPayTime string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// 失败类型
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 失败原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 实际支付金额
	ActualPayPrice int64 `json:"actual_pay_price,omitempty" xml:"actual_pay_price,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}

var poolBtripFlightPayOrderRs = sync.Pool{
	New: func() any {
		return new(BtripFlightPayOrderRs)
	},
}

// GetBtripFlightPayOrderRs() 从对象池中获取BtripFlightPayOrderRs
func GetBtripFlightPayOrderRs() *BtripFlightPayOrderRs {
	return poolBtripFlightPayOrderRs.Get().(*BtripFlightPayOrderRs)
}

// ReleaseBtripFlightPayOrderRs 释放BtripFlightPayOrderRs
func ReleaseBtripFlightPayOrderRs(v *BtripFlightPayOrderRs) {
	v.AlipayTradeNo = ""
	v.LastPayTime = ""
	v.FailCode = ""
	v.FailReason = ""
	v.ActualPayPrice = 0
	v.PayStatus = 0
	poolBtripFlightPayOrderRs.Put(v)
}
