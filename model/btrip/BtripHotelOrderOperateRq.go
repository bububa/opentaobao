package btrip

import (
	"sync"
)

// BtripHotelOrderOperateRq 结构体
type BtripHotelOrderOperateRq struct {
	// 分销商订单id
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销商子渠道，通常是商旅企业id
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 供应商标识
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 商旅订单id
	BtripOrderId int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
}

var poolBtripHotelOrderOperateRq = sync.Pool{
	New: func() any {
		return new(BtripHotelOrderOperateRq)
	},
}

// GetBtripHotelOrderOperateRq() 从对象池中获取BtripHotelOrderOperateRq
func GetBtripHotelOrderOperateRq() *BtripHotelOrderOperateRq {
	return poolBtripHotelOrderOperateRq.Get().(*BtripHotelOrderOperateRq)
}

// ReleaseBtripHotelOrderOperateRq 释放BtripHotelOrderOperateRq
func ReleaseBtripHotelOrderOperateRq(v *BtripHotelOrderOperateRq) {
	v.DisOrderId = ""
	v.SubChannel = ""
	v.SupplierCode = ""
	v.BtripOrderId = 0
	poolBtripHotelOrderOperateRq.Put(v)
}
