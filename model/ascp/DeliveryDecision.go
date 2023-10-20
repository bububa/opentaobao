package ascp

import (
	"sync"
)

// DeliveryDecision 结构体
type DeliveryDecision struct {
	// ERP发货单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 订单发货地，省份
	SendProvince string `json:"send_province,omitempty" xml:"send_province,omitempty"`
	// 订单发货地，所在城市
	SendCity string `json:"send_city,omitempty" xml:"send_city,omitempty"`
	// 订单发货地，所在地区
	SendDistrict string `json:"send_district,omitempty" xml:"send_district,omitempty"`
	// 订单发货地，街道地址
	SendTown string `json:"send_town,omitempty" xml:"send_town,omitempty"`
	// 订单发货地地址编码（先识别编码，如果识别失败，解析地址）
	SendDivisionCode string `json:"send_division_code,omitempty" xml:"send_division_code,omitempty"`
	// 主交易单号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 子交易单号
	SubTradeId string `json:"sub_trade_id,omitempty" xml:"sub_trade_id,omitempty"`
	// 订单类型 枚举： FENXIAO=分销订单 ，默认店铺零售订单。
	OrderFlag string `json:"order_flag,omitempty" xml:"order_flag,omitempty"`
	// 订单收货地，省份
	ReceiveProvince string `json:"receive_province,omitempty" xml:"receive_province,omitempty"`
	// 订单收货地，所在城市
	ReceiveCity string `json:"receive_city,omitempty" xml:"receive_city,omitempty"`
	// 订单收货地，所在地区
	ReceiveDistrict string `json:"receive_district,omitempty" xml:"receive_district,omitempty"`
	// 订单收货地，街道地址
	ReceiveTown string `json:"receive_town,omitempty" xml:"receive_town,omitempty"`
	// 订单收货地地址编码（先识别编码，如果识别失败，解析地址）
	ReceiveDivisionCode string `json:"receive_division_code,omitempty" xml:"receive_division_code,omitempty"`
}

var poolDeliveryDecision = sync.Pool{
	New: func() any {
		return new(DeliveryDecision)
	},
}

// GetDeliveryDecision() 从对象池中获取DeliveryDecision
func GetDeliveryDecision() *DeliveryDecision {
	return poolDeliveryDecision.Get().(*DeliveryDecision)
}

// ReleaseDeliveryDecision 释放DeliveryDecision
func ReleaseDeliveryDecision(v *DeliveryDecision) {
	v.OrderCode = ""
	v.SendProvince = ""
	v.SendCity = ""
	v.SendDistrict = ""
	v.SendTown = ""
	v.SendDivisionCode = ""
	v.TradeId = ""
	v.SubTradeId = ""
	v.OrderFlag = ""
	v.ReceiveProvince = ""
	v.ReceiveCity = ""
	v.ReceiveDistrict = ""
	v.ReceiveTown = ""
	v.ReceiveDivisionCode = ""
	poolDeliveryDecision.Put(v)
}
