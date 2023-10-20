package axintrade

import (
	"sync"
)

// HotelOrderRefundResApiDto 结构体
type HotelOrderRefundResApiDto struct {
	// 退款单号
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}

var poolHotelOrderRefundResApiDto = sync.Pool{
	New: func() any {
		return new(HotelOrderRefundResApiDto)
	},
}

// GetHotelOrderRefundResApiDto() 从对象池中获取HotelOrderRefundResApiDto
func GetHotelOrderRefundResApiDto() *HotelOrderRefundResApiDto {
	return poolHotelOrderRefundResApiDto.Get().(*HotelOrderRefundResApiDto)
}

// ReleaseHotelOrderRefundResApiDto 释放HotelOrderRefundResApiDto
func ReleaseHotelOrderRefundResApiDto(v *HotelOrderRefundResApiDto) {
	v.RefundOrderId = 0
	v.RefundFee = 0
	poolHotelOrderRefundResApiDto.Put(v)
}
