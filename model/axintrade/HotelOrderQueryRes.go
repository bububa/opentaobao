package axintrade

import (
	"sync"
)

// HotelOrderQueryRes 结构体
type HotelOrderQueryRes struct {
	// 每日价格信息
	DailyInfoList []DailyInfo `json:"daily_info_list,omitempty" xml:"daily_info_list>daily_info,omitempty"`
	// 订单状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 飞猪订单号
	FliggyOrderId string `json:"fliggy_order_id,omitempty" xml:"fliggy_order_id,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 酒店信息
	HotelInfo *HotelInfo `json:"hotel_info,omitempty" xml:"hotel_info,omitempty"`
	// 房间信息
	RoomInfo *RoomInfo `json:"room_info,omitempty" xml:"room_info,omitempty"`
	// 订单履约信息
	OrderFulfillInfo *OrderFulfillInfo `json:"order_fulfill_info,omitempty" xml:"order_fulfill_info,omitempty"`
	// 实际支付总金额，单位为分
	ActualTotalFee int64 `json:"actual_total_fee,omitempty" xml:"actual_total_fee,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 采购单id
	PurchaseOrderId int64 `json:"purchase_order_id,omitempty" xml:"purchase_order_id,omitempty"`
	// 退款金额，单位为分（产生退款时不为空）
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 支付id
	PayId int64 `json:"pay_id,omitempty" xml:"pay_id,omitempty"`
	// 汇率
	ExchangeRate float64 `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 实际支付原币种金额
	OriginActualTotalFee int64 `json:"origin_actual_total_fee,omitempty" xml:"origin_actual_total_fee,omitempty"`
}

var poolHotelOrderQueryRes = sync.Pool{
	New: func() any {
		return new(HotelOrderQueryRes)
	},
}

// GetHotelOrderQueryRes() 从对象池中获取HotelOrderQueryRes
func GetHotelOrderQueryRes() *HotelOrderQueryRes {
	return poolHotelOrderQueryRes.Get().(*HotelOrderQueryRes)
}

// ReleaseHotelOrderQueryRes 释放HotelOrderQueryRes
func ReleaseHotelOrderQueryRes(v *HotelOrderQueryRes) {
	v.DailyInfoList = v.DailyInfoList[:0]
	v.StatusDesc = ""
	v.OuterOrderId = ""
	v.FliggyOrderId = ""
	v.CurrencyCode = ""
	v.HotelInfo = nil
	v.RoomInfo = nil
	v.OrderFulfillInfo = nil
	v.ActualTotalFee = 0
	v.Status = 0
	v.PurchaseOrderId = 0
	v.RefundFee = 0
	v.PayId = 0
	v.ExchangeRate = 0
	v.OriginActualTotalFee = 0
	poolHotelOrderQueryRes.Put(v)
}
