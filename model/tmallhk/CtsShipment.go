package tmallhk

import (
	"sync"
)

// CtsShipment 结构体
type CtsShipment struct {
	// 报关开始时间
	Begin string `json:"begin,omitempty" xml:"begin,omitempty"`
	// 报关结束时间
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 报关单号
	ShipmentNo string `json:"shipment_no,omitempty" xml:"shipment_no,omitempty"`
}

var poolCtsShipment = sync.Pool{
	New: func() any {
		return new(CtsShipment)
	},
}

// GetCtsShipment() 从对象池中获取CtsShipment
func GetCtsShipment() *CtsShipment {
	return poolCtsShipment.Get().(*CtsShipment)
}

// ReleaseCtsShipment 释放CtsShipment
func ReleaseCtsShipment(v *CtsShipment) {
	v.Begin = ""
	v.End = ""
	v.ShipmentNo = ""
	poolCtsShipment.Put(v)
}
