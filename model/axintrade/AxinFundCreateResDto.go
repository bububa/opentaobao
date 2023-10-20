package axintrade

import (
	"sync"
)

// AxinFundCreateResDto 结构体
type AxinFundCreateResDto struct {
	// 支付宝返回的拼接串
	AlipayBody string `json:"alipay_body,omitempty" xml:"alipay_body,omitempty"`
	// 阿信支付宝账号appID
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 资金单ID
	FundId string `json:"fund_id,omitempty" xml:"fund_id,omitempty"`
	// 商户原始订单号，最大长度限制32位
	MerchantOrderNo string `json:"merchant_order_no,omitempty" xml:"merchant_order_no,omitempty"`
	// 商户订单号，由商家自定义
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 收款支付宝账号对应的支付宝唯一用户号
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 该笔订单的资金总额，单位为RMB-Yuan
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 该交易在支付宝系统中的交易流水号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
}

var poolAxinFundCreateResDto = sync.Pool{
	New: func() any {
		return new(AxinFundCreateResDto)
	},
}

// GetAxinFundCreateResDto() 从对象池中获取AxinFundCreateResDto
func GetAxinFundCreateResDto() *AxinFundCreateResDto {
	return poolAxinFundCreateResDto.Get().(*AxinFundCreateResDto)
}

// ReleaseAxinFundCreateResDto 释放AxinFundCreateResDto
func ReleaseAxinFundCreateResDto(v *AxinFundCreateResDto) {
	v.AlipayBody = ""
	v.AppId = ""
	v.FundId = ""
	v.MerchantOrderNo = ""
	v.OutTradeNo = ""
	v.SellerId = ""
	v.TotalAmount = ""
	v.TradeNo = ""
	poolAxinFundCreateResDto.Put(v)
}
