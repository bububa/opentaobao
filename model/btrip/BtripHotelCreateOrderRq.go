package btrip

import (
	"sync"
)

// BtripHotelCreateOrderRq 结构体
type BtripHotelCreateOrderRq struct {
	// 预订人在分销商平台的用户昵称
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 预订人在分销商平台的唯一用户标识
	BuyerUniqueKey string `json:"buyer_unique_key,omitempty" xml:"buyer_unique_key,omitempty"`
	// 1
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 1
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 创单Key值，试单结果返回
	CreateOrderKey string `json:"create_order_key,omitempty" xml:"create_order_key,omitempty"`
	// 入住人信息
	Customers string `json:"customers,omitempty" xml:"customers,omitempty"`
	// 分销平台订单id
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 最早到店时间
	EarliestArrivalTime string `json:"earliest_arrival_time,omitempty" xml:"earliest_arrival_time,omitempty"`
	// 最晚到店时间
	LatestArrivalTime string `json:"latest_arrival_time,omitempty" xml:"latest_arrival_time,omitempty"`
	// 分销子渠道，商旅企业id
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 供应商标识码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 订单联系人信息
	HotelContact *BtripHotelContactDto `json:"hotel_contact,omitempty" xml:"hotel_contact,omitempty"`
	// 商旅商品唯一标识
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 每间房成人数
	NumberOfAdultsPerRoom int64 `json:"number_of_adults_per_room,omitempty" xml:"number_of_adults_per_room,omitempty"`
	// 销售计划id
	RatePlanId int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// 预订房间数
	RoomNum int64 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// 总优惠金额，注:(验价时如有优惠信息，则这个字段必传)
	TotalPromotion int64 `json:"total_promotion,omitempty" xml:"total_promotion,omitempty"`
	// 总房价
	TotalRoomPrice int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 订单总价中企业支付部分
	CorpPayPrice int64 `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// 订单总价中个人支付部分
	PersonPayPrice int64 `json:"person_pay_price,omitempty" xml:"person_pay_price,omitempty"`
}

var poolBtripHotelCreateOrderRq = sync.Pool{
	New: func() any {
		return new(BtripHotelCreateOrderRq)
	},
}

// GetBtripHotelCreateOrderRq() 从对象池中获取BtripHotelCreateOrderRq
func GetBtripHotelCreateOrderRq() *BtripHotelCreateOrderRq {
	return poolBtripHotelCreateOrderRq.Get().(*BtripHotelCreateOrderRq)
}

// ReleaseBtripHotelCreateOrderRq 释放BtripHotelCreateOrderRq
func ReleaseBtripHotelCreateOrderRq(v *BtripHotelCreateOrderRq) {
	v.BuyerName = ""
	v.BuyerUniqueKey = ""
	v.CheckIn = ""
	v.CheckOut = ""
	v.CreateOrderKey = ""
	v.Customers = ""
	v.DisOrderId = ""
	v.EarliestArrivalTime = ""
	v.LatestArrivalTime = ""
	v.SubChannel = ""
	v.SupplierCode = ""
	v.HotelContact = nil
	v.ItemId = 0
	v.NumberOfAdultsPerRoom = 0
	v.RatePlanId = 0
	v.RoomNum = 0
	v.TotalPromotion = 0
	v.TotalRoomPrice = 0
	v.CorpPayPrice = 0
	v.PersonPayPrice = 0
	poolBtripHotelCreateOrderRq.Put(v)
}
