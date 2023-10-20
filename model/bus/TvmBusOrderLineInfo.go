package bus

import (
	"sync"
)

// TvmBusOrderLineInfo 结构体
type TvmBusOrderLineInfo struct {
	// passengers
	Passengers []TvmPassengerVo `json:"passengers,omitempty" xml:"passengers>tvm_passenger_vo,omitempty"`
	// 退款信息
	Refunds []TvmRefundApply `json:"refunds,omitempty" xml:"refunds>tvm_refund_apply,omitempty"`
	// 代理商订单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 阿里支付交易号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 阿里订单编号
	AlitripOrderId string `json:"alitrip_order_id,omitempty" xml:"alitrip_order_id,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 出票时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 实际支付方式:取值范围: ALIPAY （飞猪渠道）; WECHAT（微信）
	RealPayMode string `json:"real_pay_mode,omitempty" xml:"real_pay_mode,omitempty"`
	// 买家用户信息唯一标识
	BuyerInfoUniqueKey string `json:"buyer_info_unique_key,omitempty" xml:"buyer_info_unique_key,omitempty"`
	// 阿里tp订单号
	AlitripTpOrderId string `json:"alitrip_tp_order_id,omitempty" xml:"alitrip_tp_order_id,omitempty"`
	// CREATED_NO_PAY(10, 已创建，待支付),PAYED_NO_NOTIFY(20, 已支付，待通知代理商),PAYED_AND_NOTIFY(30, 已支付，待代理商出票),BOOKED_AND_CONFIRM（50, 已出票,已确认）,BOOK_FAILED(-1, 出票失败),PAY_TIME_OUT(-2, 支付超时),BOOK_TIME_OUT(-3, 出票超时),CANCEL(-4, 订单取消),CLOSE(-5, 订单关闭)
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 票数
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// totalPrice
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// tvmBusLineInfo
	TvmBusLineInfo *TvmBusLineInfo `json:"tvm_bus_line_info,omitempty" xml:"tvm_bus_line_info,omitempty"`
}

var poolTvmBusOrderLineInfo = sync.Pool{
	New: func() any {
		return new(TvmBusOrderLineInfo)
	},
}

// GetTvmBusOrderLineInfo() 从对象池中获取TvmBusOrderLineInfo
func GetTvmBusOrderLineInfo() *TvmBusOrderLineInfo {
	return poolTvmBusOrderLineInfo.Get().(*TvmBusOrderLineInfo)
}

// ReleaseTvmBusOrderLineInfo 释放TvmBusOrderLineInfo
func ReleaseTvmBusOrderLineInfo(v *TvmBusOrderLineInfo) {
	v.Passengers = v.Passengers[:0]
	v.Refunds = v.Refunds[:0]
	v.AgentOrderId = ""
	v.AlipayTradeNo = ""
	v.AlitripOrderId = ""
	v.GmtCreate = ""
	v.IssueTime = ""
	v.PayTime = ""
	v.RealPayMode = ""
	v.BuyerInfoUniqueKey = ""
	v.AlitripTpOrderId = ""
	v.OrderStatus = 0
	v.TicketCount = 0
	v.TotalPrice = 0
	v.TvmBusLineInfo = nil
	poolTvmBusOrderLineInfo.Put(v)
}
