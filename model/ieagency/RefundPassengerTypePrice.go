package ieagency

import (
	"sync"
)

// RefundPassengerTypePrice 结构体
type RefundPassengerTypePrice struct {
	// 机票价格信息
	FlightPrice *RefundFlightPrice `json:"flight_price,omitempty" xml:"flight_price,omitempty"`
	// 乘机人类型(Adult(0, &#34;成人&#34;),     Child(1, &#34;儿童&#34;),     StudentAbroad(2, &#34;留学生&#34;),     Infant(3, &#34;婴儿&#34;)
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
}

var poolRefundPassengerTypePrice = sync.Pool{
	New: func() any {
		return new(RefundPassengerTypePrice)
	},
}

// GetRefundPassengerTypePrice() 从对象池中获取RefundPassengerTypePrice
func GetRefundPassengerTypePrice() *RefundPassengerTypePrice {
	return poolRefundPassengerTypePrice.Get().(*RefundPassengerTypePrice)
}

// ReleaseRefundPassengerTypePrice 释放RefundPassengerTypePrice
func ReleaseRefundPassengerTypePrice(v *RefundPassengerTypePrice) {
	v.FlightPrice = nil
	v.PassengerType = 0
	poolRefundPassengerTypePrice.Put(v)
}
