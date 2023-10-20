package alitripmerchant

import (
	"sync"
)

// VoucherOrderDetailVo 结构体
type VoucherOrderDetailVo struct {
	// 订单入住宾客信息
	GuestByRoom []GuestByRoomDto `json:"guest_by_room,omitempty" xml:"guest_by_room>guest_by_room_dto,omitempty"`
	// 套餐订单详情信息
	VoucherOrder *VoucherOrderVo `json:"voucher_order,omitempty" xml:"voucher_order,omitempty"`
	// 酒店详情信息
	VoucherHotel *VoucherHotelVo `json:"voucher_hotel,omitempty" xml:"voucher_hotel,omitempty"`
	// 费用明细对象
	PriceDetailDto *VoucherHotelVo `json:"price_detail_dto,omitempty" xml:"price_detail_dto,omitempty"`
	// 房间信息详情
	OrderRoomDetail *PriceDetailDto `json:"order_room_detail,omitempty" xml:"order_room_detail,omitempty"`
	// 权益商品相关展示信息
	DerbyVoucherInfo *DerbyVoucherInfo `json:"derby_voucher_info,omitempty" xml:"derby_voucher_info,omitempty"`
}

var poolVoucherOrderDetailVo = sync.Pool{
	New: func() any {
		return new(VoucherOrderDetailVo)
	},
}

// GetVoucherOrderDetailVo() 从对象池中获取VoucherOrderDetailVo
func GetVoucherOrderDetailVo() *VoucherOrderDetailVo {
	return poolVoucherOrderDetailVo.Get().(*VoucherOrderDetailVo)
}

// ReleaseVoucherOrderDetailVo 释放VoucherOrderDetailVo
func ReleaseVoucherOrderDetailVo(v *VoucherOrderDetailVo) {
	v.GuestByRoom = v.GuestByRoom[:0]
	v.VoucherOrder = nil
	v.VoucherHotel = nil
	v.PriceDetailDto = nil
	v.OrderRoomDetail = nil
	v.DerbyVoucherInfo = nil
	poolVoucherOrderDetailVo.Put(v)
}
