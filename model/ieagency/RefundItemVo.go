package ieagency

import (
	"sync"
)

// RefundItemVo 结构体
type RefundItemVo struct {
	// 乘机人类型价格信息
	PassengerTypePrices []RefundPassengerTypePrice `json:"passenger_type_prices,omitempty" xml:"passenger_type_prices>refund_passenger_type_price,omitempty"`
	// 行程信息
	RefundItineraryFlights []RefundItineraryVo `json:"refund_itinerary_flights,omitempty" xml:"refund_itinerary_flights>refund_itinerary_vo,omitempty"`
}

var poolRefundItemVo = sync.Pool{
	New: func() any {
		return new(RefundItemVo)
	},
}

// GetRefundItemVo() 从对象池中获取RefundItemVo
func GetRefundItemVo() *RefundItemVo {
	return poolRefundItemVo.Get().(*RefundItemVo)
}

// ReleaseRefundItemVo 释放RefundItemVo
func ReleaseRefundItemVo(v *RefundItemVo) {
	v.PassengerTypePrices = v.PassengerTypePrices[:0]
	v.RefundItineraryFlights = v.RefundItineraryFlights[:0]
	poolRefundItemVo.Put(v)
}
