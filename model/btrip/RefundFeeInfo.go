package btrip

import (
	"sync"
)

// RefundFeeInfo 结构体
type RefundFeeInfo struct {
	// 支付宝交易订单号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 退款状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 退票金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退票手续费
	RefundPrice int64 `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
}

var poolRefundFeeInfo = sync.Pool{
	New: func() any {
		return new(RefundFeeInfo)
	},
}

// GetRefundFeeInfo() 从对象池中获取RefundFeeInfo
func GetRefundFeeInfo() *RefundFeeInfo {
	return poolRefundFeeInfo.Get().(*RefundFeeInfo)
}

// ReleaseRefundFeeInfo 释放RefundFeeInfo
func ReleaseRefundFeeInfo(v *RefundFeeInfo) {
	v.AlipayTradeNo = ""
	v.Status = ""
	v.RefundFee = 0
	v.RefundPrice = 0
	poolRefundFeeInfo.Put(v)
}
