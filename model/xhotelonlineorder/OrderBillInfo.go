package xhotelonlineorder

import (
	"sync"
)

// OrderBillInfo 结构体
type OrderBillInfo struct {
	// 离店日期
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 入住日期
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 房间号
	RoomNo string `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// 完整的身份证号
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 客人姓名
	GuestName string `json:"guest_name,omitempty" xml:"guest_name,omitempty"`
	// 外部酒店代码
	HotelCode string `json:"hotel_code,omitempty" xml:"hotel_code,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 请求id (同入参)
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 未结账房费
	NoCheckoutPrice int64 `json:"no_checkout_price,omitempty" xml:"no_checkout_price,omitempty"`
	// 已结账房费
	CheckoutPrice int64 `json:"checkout_price,omitempty" xml:"checkout_price,omitempty"`
	// 已结账总费用
	CheckoutTotalFee int64 `json:"checkout_total_fee,omitempty" xml:"checkout_total_fee,omitempty"`
	// 未结账总费用
	NoCheckoutTotalFee int64 `json:"no_checkout_total_fee,omitempty" xml:"no_checkout_total_fee,omitempty"`
	// 淘宝订单号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 未结账杂费
	NoCheckoutOtherFee int64 `json:"no_checkout_other_fee,omitempty" xml:"no_checkout_other_fee,omitempty"`
	// 已结账杂费
	CheckoutOtherFee int64 `json:"checkout_other_fee,omitempty" xml:"checkout_other_fee,omitempty"`
	// 杂费明细列表
	OtherFeeDetail *OtherFeeDetail `json:"other_fee_detail,omitempty" xml:"other_fee_detail,omitempty"`
	// 每日房费类表
	DailyRoomFee *DailyRoomFee `json:"daily_room_fee,omitempty" xml:"daily_room_fee,omitempty"`
}

var poolOrderBillInfo = sync.Pool{
	New: func() any {
		return new(OrderBillInfo)
	},
}

// GetOrderBillInfo() 从对象池中获取OrderBillInfo
func GetOrderBillInfo() *OrderBillInfo {
	return poolOrderBillInfo.Get().(*OrderBillInfo)
}

// ReleaseOrderBillInfo 释放OrderBillInfo
func ReleaseOrderBillInfo(v *OrderBillInfo) {
	v.CheckOutDate = ""
	v.CheckInDate = ""
	v.RoomNo = ""
	v.IdNumber = ""
	v.GuestName = ""
	v.HotelCode = ""
	v.OutOrderId = ""
	v.RequestId = ""
	v.Remark = ""
	v.NoCheckoutPrice = 0
	v.CheckoutPrice = 0
	v.CheckoutTotalFee = 0
	v.NoCheckoutTotalFee = 0
	v.Tid = 0
	v.NoCheckoutOtherFee = 0
	v.CheckoutOtherFee = 0
	v.OtherFeeDetail = nil
	v.DailyRoomFee = nil
	poolOrderBillInfo.Put(v)
}
