package traveltrade

import (
	"sync"
)

// PayInfo 结构体
type PayInfo struct {
	// 支付宝交易号
	AlipayNo string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	// 订单付款时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 分阶段付款的订单状态，目前有三返回状态FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付)
	StepTradeStatus string `json:"step_trade_status,omitempty" xml:"step_trade_status,omitempty"`
	// 系统优惠金额（如打折，VIP，满就送等），精确到2位小数，单位：分。
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:分
	ReceivedPayment int64 `json:"received_payment,omitempty" xml:"received_payment,omitempty"`
	// 分阶段付款的已付金额
	StepPaidFee int64 `json:"step_paid_fee,omitempty" xml:"step_paid_fee,omitempty"`
}

var poolPayInfo = sync.Pool{
	New: func() any {
		return new(PayInfo)
	},
}

// GetPayInfo() 从对象池中获取PayInfo
func GetPayInfo() *PayInfo {
	return poolPayInfo.Get().(*PayInfo)
}

// ReleasePayInfo 释放PayInfo
func ReleasePayInfo(v *PayInfo) {
	v.AlipayNo = ""
	v.PayTime = ""
	v.StepTradeStatus = ""
	v.DiscountFee = 0
	v.ReceivedPayment = 0
	v.StepPaidFee = 0
	poolPayInfo.Put(v)
}
