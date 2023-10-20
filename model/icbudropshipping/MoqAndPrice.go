package icbudropshipping

import (
	"sync"
)

// MoqAndPrice 结构体
type MoqAndPrice struct {
	// min order quantity
	MinOrderQuantity string `json:"min_order_quantity,omitempty" xml:"min_order_quantity,omitempty"`
	// min order quantity unit
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// min order quantity  delivery period
	MoqDeliveryPeriod int64 `json:"moq_delivery_period,omitempty" xml:"moq_delivery_period,omitempty"`
	// min order quantity unit price
	MoqUnitPrice float64 `json:"moq_unit_price,omitempty" xml:"moq_unit_price,omitempty"`
}

var poolMoqAndPrice = sync.Pool{
	New: func() any {
		return new(MoqAndPrice)
	},
}

// GetMoqAndPrice() 从对象池中获取MoqAndPrice
func GetMoqAndPrice() *MoqAndPrice {
	return poolMoqAndPrice.Get().(*MoqAndPrice)
}

// ReleaseMoqAndPrice 释放MoqAndPrice
func ReleaseMoqAndPrice(v *MoqAndPrice) {
	v.MinOrderQuantity = ""
	v.Unit = ""
	v.MoqDeliveryPeriod = 0
	v.MoqUnitPrice = 0
	poolMoqAndPrice.Put(v)
}
