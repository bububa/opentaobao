package btrip

import (
	"sync"
)

// PassengerFlightShoppingItemRs 结构体
type PassengerFlightShoppingItemRs struct {
	// 乘客类型
	PassengerType string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 机票报价
	ShoppingItem *FlightShoppingItemRs `json:"shopping_item,omitempty" xml:"shopping_item,omitempty"`
}

var poolPassengerFlightShoppingItemRs = sync.Pool{
	New: func() any {
		return new(PassengerFlightShoppingItemRs)
	},
}

// GetPassengerFlightShoppingItemRs() 从对象池中获取PassengerFlightShoppingItemRs
func GetPassengerFlightShoppingItemRs() *PassengerFlightShoppingItemRs {
	return poolPassengerFlightShoppingItemRs.Get().(*PassengerFlightShoppingItemRs)
}

// ReleasePassengerFlightShoppingItemRs 释放PassengerFlightShoppingItemRs
func ReleasePassengerFlightShoppingItemRs(v *PassengerFlightShoppingItemRs) {
	v.PassengerType = ""
	v.ShoppingItem = nil
	poolPassengerFlightShoppingItemRs.Put(v)
}
