package xhotelonlineorder

import (
	"sync"
)

// HotelOrderSellerDo 结构体
type HotelOrderSellerDo struct {
	// 支付类型
	Payment *OrderSellerPaymentDo `json:"payment,omitempty" xml:"payment,omitempty"`
	// 卖家结账单信息
	SettlePayment *OrderSettleDo `json:"settle_payment,omitempty" xml:"settle_payment,omitempty"`
}

var poolHotelOrderSellerDo = sync.Pool{
	New: func() any {
		return new(HotelOrderSellerDo)
	},
}

// GetHotelOrderSellerDo() 从对象池中获取HotelOrderSellerDo
func GetHotelOrderSellerDo() *HotelOrderSellerDo {
	return poolHotelOrderSellerDo.Get().(*HotelOrderSellerDo)
}

// ReleaseHotelOrderSellerDo 释放HotelOrderSellerDo
func ReleaseHotelOrderSellerDo(v *HotelOrderSellerDo) {
	v.Payment = nil
	v.SettlePayment = nil
	poolHotelOrderSellerDo.Put(v)
}
