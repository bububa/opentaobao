package alitripmerchant

import (
	"sync"
)

// OrderDto 结构体
type OrderDto struct {
	// 床型名称
	BedName string `json:"bed_name,omitempty" xml:"bed_name,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 用户支付总价
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 预订日期
	BookDate string `json:"book_date,omitempty" xml:"book_date,omitempty"`
	// 入住人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 离店时间
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 房型照片
	RoomPhotoUrl string `json:"room_photo_url,omitempty" xml:"room_photo_url,omitempty"`
	// 酒店外部id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 下单类型
	PlaceOrderType string `json:"place_order_type,omitempty" xml:"place_order_type,omitempty"`
	// 通用券id
	UniversalCouponId string `json:"universal_coupon_id,omitempty" xml:"universal_coupon_id,omitempty"`
	// 差价
	Spread string `json:"spread,omitempty" xml:"spread,omitempty"`
	// 代金券名称
	VoucherName string `json:"voucher_name,omitempty" xml:"voucher_name,omitempty"`
	// 支付剩余时间
	PayRemainTime int64 `json:"pay_remain_time,omitempty" xml:"pay_remain_time,omitempty"`
	// 房间数量
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 酒店房型id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 成人数量
	AdultNumber int64 `json:"adult_number,omitempty" xml:"adult_number,omitempty"`
	// 儿童数量
	ChildrenNumber int64 `json:"children_number,omitempty" xml:"children_number,omitempty"`
	// 外币信息
	ForeignCurrency *ForeignCurrencyInfo `json:"foreign_currency,omitempty" xml:"foreign_currency,omitempty"`
	// 是否为外币支付
	ForeignCurrencyPayment bool `json:"foreign_currency_payment,omitempty" xml:"foreign_currency_payment,omitempty"`
}

var poolOrderDto = sync.Pool{
	New: func() any {
		return new(OrderDto)
	},
}

// GetOrderDto() 从对象池中获取OrderDto
func GetOrderDto() *OrderDto {
	return poolOrderDto.Get().(*OrderDto)
}

// ReleaseOrderDto 释放OrderDto
func ReleaseOrderDto(v *OrderDto) {
	v.BedName = ""
	v.Currency = ""
	v.OrderStatus = ""
	v.TotalAmount = ""
	v.BookDate = ""
	v.ContactName = ""
	v.CheckOutDate = ""
	v.CheckInDate = ""
	v.RoomName = ""
	v.HotelName = ""
	v.OrderCode = ""
	v.RoomPhotoUrl = ""
	v.HotelId = ""
	v.PlaceOrderType = ""
	v.UniversalCouponId = ""
	v.Spread = ""
	v.VoucherName = ""
	v.PayRemainTime = 0
	v.RoomNumber = 0
	v.Shid = 0
	v.AdultNumber = 0
	v.ChildrenNumber = 0
	v.ForeignCurrency = nil
	v.ForeignCurrencyPayment = false
	poolOrderDto.Put(v)
}
