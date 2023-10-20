package btrip

import (
	"sync"
)

// BtripFlightModifyDetailInfoRs 结构体
type BtripFlightModifyDetailInfoRs struct {
	// 航班信息
	FlightInfoList []OrderFlightInfo `json:"flight_info_list,omitempty" xml:"flight_info_list>order_flight_info,omitempty"`
	// 乘机人信息
	TravelerInfoList []OrderTravelerInfo `json:"traveler_info_list,omitempty" xml:"traveler_info_list>order_traveler_info,omitempty"`
	// 交易流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 最迟支付时间
	LastPayTime string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// 支付状态
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 分销外部改签单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 结算费用
	SettlePrice int64 `json:"settle_price,omitempty" xml:"settle_price,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 总改签费用
	TotalChangePrice int64 `json:"total_change_price,omitempty" xml:"total_change_price,omitempty"`
	// 总费用
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 总升舱费用
	TotalUpgradePrice int64 `json:"total_upgrade_price,omitempty" xml:"total_upgrade_price,omitempty"`
	// 商旅改签单号
	BtripSubOrderId int64 `json:"btrip_sub_order_id,omitempty" xml:"btrip_sub_order_id,omitempty"`
	// 商旅订单号
	BtripOrderId int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// 结算类型
	SettleType int64 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
}

var poolBtripFlightModifyDetailInfoRs = sync.Pool{
	New: func() any {
		return new(BtripFlightModifyDetailInfoRs)
	},
}

// GetBtripFlightModifyDetailInfoRs() 从对象池中获取BtripFlightModifyDetailInfoRs
func GetBtripFlightModifyDetailInfoRs() *BtripFlightModifyDetailInfoRs {
	return poolBtripFlightModifyDetailInfoRs.Get().(*BtripFlightModifyDetailInfoRs)
}

// ReleaseBtripFlightModifyDetailInfoRs 释放BtripFlightModifyDetailInfoRs
func ReleaseBtripFlightModifyDetailInfoRs(v *BtripFlightModifyDetailInfoRs) {
	v.FlightInfoList = v.FlightInfoList[:0]
	v.TravelerInfoList = v.TravelerInfoList[:0]
	v.AlipayTradeNo = ""
	v.Extra = ""
	v.LastPayTime = ""
	v.PayStatus = ""
	v.PayTime = ""
	v.DisSubOrderId = ""
	v.DisOrderId = ""
	v.SettlePrice = 0
	v.Status = 0
	v.TotalChangePrice = 0
	v.TotalPrice = 0
	v.TotalUpgradePrice = 0
	v.BtripSubOrderId = 0
	v.BtripOrderId = 0
	v.SettleType = 0
	poolBtripFlightModifyDetailInfoRs.Put(v)
}
