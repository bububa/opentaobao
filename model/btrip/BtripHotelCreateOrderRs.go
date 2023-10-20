package btrip

import (
	"sync"
)

// BtripHotelCreateOrderRs 结构体
type BtripHotelCreateOrderRs struct {
	// 供应商订单id
	SupplierOrderId string `json:"supplier_order_id,omitempty" xml:"supplier_order_id,omitempty"`
	// 商旅订单id
	BtripOrderId int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
}

var poolBtripHotelCreateOrderRs = sync.Pool{
	New: func() any {
		return new(BtripHotelCreateOrderRs)
	},
}

// GetBtripHotelCreateOrderRs() 从对象池中获取BtripHotelCreateOrderRs
func GetBtripHotelCreateOrderRs() *BtripHotelCreateOrderRs {
	return poolBtripHotelCreateOrderRs.Get().(*BtripHotelCreateOrderRs)
}

// ReleaseBtripHotelCreateOrderRs 释放BtripHotelCreateOrderRs
func ReleaseBtripHotelCreateOrderRs(v *BtripHotelCreateOrderRs) {
	v.SupplierOrderId = ""
	v.BtripOrderId = 0
	poolBtripHotelCreateOrderRs.Put(v)
}
