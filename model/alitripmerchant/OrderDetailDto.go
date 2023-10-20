package alitripmerchant

import (
	"sync"
)

// OrderDetailDto 结构体
type OrderDetailDto struct {
	// 每个房间入住人信息
	GuestByRoomDtos []GuestByRoomDto `json:"guest_by_room_dtos,omitempty" xml:"guest_by_room_dtos>guest_by_room_dto,omitempty"`
	// 酒店电话
	HotelPhone string `json:"hotel_phone,omitempty" xml:"hotel_phone,omitempty"`
	// 早餐类别
	BreakfastType string `json:"breakfast_type,omitempty" xml:"breakfast_type,omitempty"`
	// 房间图片
	RoomPicture string `json:"room_picture,omitempty" xml:"room_picture,omitempty"`
	// 预定时间
	PaymentDate string `json:"payment_date,omitempty" xml:"payment_date,omitempty"`
	// 订单编号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lon string `json:"lon,omitempty" xml:"lon,omitempty"`
	// 入住人邮箱
	ContactEmail string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// 入住人手机
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 入住人姓
	ContactFirstName string `json:"contact_first_name,omitempty" xml:"contact_first_name,omitempty"`
	// 入住人名
	ContactLastName string `json:"contact_last_name,omitempty" xml:"contact_last_name,omitempty"`
	// 离店时间
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 价格名称
	RpName string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
	// 房间床型描述
	BedTypeDesc string `json:"bed_type_desc,omitempty" xml:"bed_type_desc,omitempty"`
	// 房间面积
	RoomArea string `json:"room_area,omitempty" xml:"room_area,omitempty"`
	// 房间名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 货币
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 总价格
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 取消描述
	CancelDec string `json:"cancel_dec,omitempty" xml:"cancel_dec,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 卖家房型id
	OuterRoomId string `json:"outer_room_id,omitempty" xml:"outer_room_id,omitempty"`
	// 租户酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 酒店地址描述
	HotelAddress string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	// 退款手续费
	RefundCostAmount string `json:"refund_cost_amount,omitempty" xml:"refund_cost_amount,omitempty"`
	// 订单状态描述
	OrderStatusDesc string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// 支付渠道
	PaymentChannel string `json:"payment_channel,omitempty" xml:"payment_channel,omitempty"`
	// 酒店预订确认参数
	PmsCode string `json:"pms_code,omitempty" xml:"pms_code,omitempty"`
	// 入住人数
	PersonNum int64 `json:"person_num,omitempty" xml:"person_num,omitempty"`
	// 入住天数
	Days int64 `json:"days,omitempty" xml:"days,omitempty"`
	// 房间数量
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 酒店房型信息
	RoomDetailDto *RoomDetailDto `json:"room_detail_dto,omitempty" xml:"room_detail_dto,omitempty"`
	// 最大入住人数
	MaxCheckInNumber int64 `json:"max_check_in_number,omitempty" xml:"max_check_in_number,omitempty"`
	// 费用明细对象
	PriceDetailDto *PriceDetailDto `json:"price_detail_dto,omitempty" xml:"price_detail_dto,omitempty"`
	// 支付类型
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 取消政策规则
	CancelRule int64 `json:"cancel_rule,omitempty" xml:"cancel_rule,omitempty"`
	// 支付剩余时间
	PayRemainTime int64 `json:"pay_remain_time,omitempty" xml:"pay_remain_time,omitempty"`
	// 卖家价格ID
	OutRateId int64 `json:"out_rate_id,omitempty" xml:"out_rate_id,omitempty"`
	// 成人总数
	AdultNumber int64 `json:"adult_number,omitempty" xml:"adult_number,omitempty"`
	// 儿童总数
	ChildrenNumber int64 `json:"children_number,omitempty" xml:"children_number,omitempty"`
}

var poolOrderDetailDto = sync.Pool{
	New: func() any {
		return new(OrderDetailDto)
	},
}

// GetOrderDetailDto() 从对象池中获取OrderDetailDto
func GetOrderDetailDto() *OrderDetailDto {
	return poolOrderDetailDto.Get().(*OrderDetailDto)
}

// ReleaseOrderDetailDto 释放OrderDetailDto
func ReleaseOrderDetailDto(v *OrderDetailDto) {
	v.GuestByRoomDtos = v.GuestByRoomDtos[:0]
	v.HotelPhone = ""
	v.BreakfastType = ""
	v.RoomPicture = ""
	v.PaymentDate = ""
	v.OrderCode = ""
	v.Lat = ""
	v.Lon = ""
	v.ContactEmail = ""
	v.ContactPhone = ""
	v.ContactFirstName = ""
	v.ContactLastName = ""
	v.CheckOutDate = ""
	v.CheckInDate = ""
	v.RpName = ""
	v.BedTypeDesc = ""
	v.RoomArea = ""
	v.RoomName = ""
	v.HotelName = ""
	v.Currency = ""
	v.TotalPrice = ""
	v.CancelDec = ""
	v.OrderStatus = ""
	v.OuterRoomId = ""
	v.HotelId = ""
	v.HotelAddress = ""
	v.RefundCostAmount = ""
	v.OrderStatusDesc = ""
	v.PaymentChannel = ""
	v.PmsCode = ""
	v.PersonNum = 0
	v.Days = 0
	v.RoomNumber = 0
	v.RoomDetailDto = nil
	v.MaxCheckInNumber = 0
	v.PriceDetailDto = nil
	v.PayType = 0
	v.CancelRule = 0
	v.PayRemainTime = 0
	v.OutRateId = 0
	v.AdultNumber = 0
	v.ChildrenNumber = 0
	poolOrderDetailDto.Put(v)
}
