package alitripmerchant

import (
	"sync"
)

// ValidateOrderParam 结构体
type ValidateOrderParam struct {
	// 多房价数据状态
	GuestByRoomDtos []GuestByRoomDto `json:"guest_by_room_dtos,omitempty" xml:"guest_by_room_dtos>guest_by_room_dto,omitempty"`
	// 总价格
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 离店时间
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 外部酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 1
	RpCode string `json:"rp_code,omitempty" xml:"rp_code,omitempty"`
	// 1
	RoomId string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// 支付类型
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 价格的rateId
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// 宝贝id
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// 价格id
	RpId int64 `json:"rp_id,omitempty" xml:"rp_id,omitempty"`
	// 飞猪外部房型id
	OutRoomId int64 `json:"out_room_id,omitempty" xml:"out_room_id,omitempty"`
	// 内部酒店id
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 入住人数
	CustomerNumbers int64 `json:"customer_numbers,omitempty" xml:"customer_numbers,omitempty"`
	// 房间数量
	RoomNum int64 `json:"room_num,omitempty" xml:"room_num,omitempty"`
}

var poolValidateOrderParam = sync.Pool{
	New: func() any {
		return new(ValidateOrderParam)
	},
}

// GetValidateOrderParam() 从对象池中获取ValidateOrderParam
func GetValidateOrderParam() *ValidateOrderParam {
	return poolValidateOrderParam.Get().(*ValidateOrderParam)
}

// ReleaseValidateOrderParam 释放ValidateOrderParam
func ReleaseValidateOrderParam(v *ValidateOrderParam) {
	v.GuestByRoomDtos = v.GuestByRoomDtos[:0]
	v.TotalPrice = ""
	v.CheckOutDate = ""
	v.CheckInDate = ""
	v.HotelId = ""
	v.RpCode = ""
	v.RoomId = ""
	v.PaymentType = 0
	v.RateId = 0
	v.Gid = 0
	v.RpId = 0
	v.OutRoomId = 0
	v.Hid = 0
	v.CustomerNumbers = 0
	v.RoomNum = 0
	poolValidateOrderParam.Put(v)
}
