package ieagency

import (
	"sync"
)

// RefundFlightPrice 结构体
type RefundFlightPrice struct {
	// 税费价格(单位：分)
	TaxPrice int64 `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
	// 机票价格(单位:分)
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

var poolRefundFlightPrice = sync.Pool{
	New: func() any {
		return new(RefundFlightPrice)
	},
}

// GetRefundFlightPrice() 从对象池中获取RefundFlightPrice
func GetRefundFlightPrice() *RefundFlightPrice {
	return poolRefundFlightPrice.Get().(*RefundFlightPrice)
}

// ReleaseRefundFlightPrice 释放RefundFlightPrice
func ReleaseRefundFlightPrice(v *RefundFlightPrice) {
	v.TaxPrice = 0
	v.TicketPrice = 0
	poolRefundFlightPrice.Put(v)
}
