package mos

import (
	"sync"
)

// OnsiteTradeQueryResponse 结构体
type OnsiteTradeQueryResponse struct {
	// 商户的实收资金渠道明细信息列表。
	FundBillList []FundBill `json:"fund_bill_list,omitempty" xml:"fund_bill_list>fund_bill,omitempty"`
	// 消费者在喵街的昵称，将用*号屏蔽部分内容。必然返回
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 消费者付款时间。在付款后才会返回，格式为 yyyy-MM-dd HH:mm:ss，时区是GMT+8
	GmtPayment string `json:"gmt_payment,omitempty" xml:"gmt_payment,omitempty"`
	// 原支付请求的商户支付流水号。必然返回
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 喵街交易流水号。必然返回
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 交易目前所处的状态。必然返回。取值：WAIT_FOR_CONFIRM，待确认，此时不能申请退款，可以撤销；WAIT_BUYER_PAY，等待消费者付款	，此时不能申请退款，可以撤销；TRADE_SUCCESS，付款成功，此时可以申请退款，可以撤销（自动申请退款）；TRADE_FINISHED，交易成功，此时可以申请退款，不可以撤销；TRADE_CLOSED，交易关闭，此时不能申请退款，不能撤销
	TradeStatus string `json:"trade_status,omitempty" xml:"trade_status,omitempty"`
	// 业务扩展参数，json格式
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 码来源，取值：MJ、M_TAO、ALIPAY
	AuthCodeSource string `json:"auth_code_source,omitempty" xml:"auth_code_source,omitempty"`
	// 消费者实付总额。单位：分
	BuyerTotalFundFee int64 `json:"buyer_total_fund_fee,omitempty" xml:"buyer_total_fund_fee,omitempty"`
	// 消费者优惠总额。消费者支付的营销渠道的总金额，单位为人民币（分）。一般来讲，是通过券／权益抵扣的总金额。消费者付款成功后，才返回此值
	BuyerTotalPromotionFee int64 `json:"buyer_total_promotion_fee,omitempty" xml:"buyer_total_promotion_fee,omitempty"`
	// 商户总营销成本。商户在本次交易中的营销总成本，单位为人民币（分）。
	StoreTotalMarketingFee int64 `json:"store_total_marketing_fee,omitempty" xml:"store_total_marketing_fee,omitempty"`
	// 商户实收总额。商户在本次交易中实收的总金额。单位：分
	StoreTotalReceivedFee int64 `json:"store_total_received_fee,omitempty" xml:"store_total_received_fee,omitempty"`
	// 本次交易支付的订单金额，单位为人民币（分）。必然返回
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
}

var poolOnsiteTradeQueryResponse = sync.Pool{
	New: func() any {
		return new(OnsiteTradeQueryResponse)
	},
}

// GetOnsiteTradeQueryResponse() 从对象池中获取OnsiteTradeQueryResponse
func GetOnsiteTradeQueryResponse() *OnsiteTradeQueryResponse {
	return poolOnsiteTradeQueryResponse.Get().(*OnsiteTradeQueryResponse)
}

// ReleaseOnsiteTradeQueryResponse 释放OnsiteTradeQueryResponse
func ReleaseOnsiteTradeQueryResponse(v *OnsiteTradeQueryResponse) {
	v.FundBillList = v.FundBillList[:0]
	v.BuyerNick = ""
	v.GmtPayment = ""
	v.OutTradeNo = ""
	v.TradeNo = ""
	v.TradeStatus = ""
	v.ExtendParams = ""
	v.AuthCodeSource = ""
	v.BuyerTotalFundFee = 0
	v.BuyerTotalPromotionFee = 0
	v.StoreTotalMarketingFee = 0
	v.StoreTotalReceivedFee = 0
	v.TotalAmount = 0
	poolOnsiteTradeQueryResponse.Put(v)
}
