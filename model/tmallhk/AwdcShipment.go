package tmallhk

import (
	"sync"
)

// AwdcShipment 结构体
type AwdcShipment struct {
	// 到达城市
	ArrivalCity string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// 抵港时间
	ArrivalDate string `json:"arrival_date,omitempty" xml:"arrival_date,omitempty"`
	// 起运城市
	DepartureCity string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// 离港时间
	DepartureDate string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
	// 报关开始时间
	DoIn string `json:"do_in,omitempty" xml:"do_in,omitempty"`
	// 报关单号
	DoNumber string `json:"do_number,omitempty" xml:"do_number,omitempty"`
	// 报关结束时间
	DoOut string `json:"do_out,omitempty" xml:"do_out,omitempty"`
	// 押运公司物流单号
	LogisticNumber string `json:"logistic_number,omitempty" xml:"logistic_number,omitempty"`
	// 航班装运订单
	ShipmentNumber string `json:"shipment_number,omitempty" xml:"shipment_number,omitempty"`
	// 押运公司
	Shipper string `json:"shipper,omitempty" xml:"shipper,omitempty"`
}

var poolAwdcShipment = sync.Pool{
	New: func() any {
		return new(AwdcShipment)
	},
}

// GetAwdcShipment() 从对象池中获取AwdcShipment
func GetAwdcShipment() *AwdcShipment {
	return poolAwdcShipment.Get().(*AwdcShipment)
}

// ReleaseAwdcShipment 释放AwdcShipment
func ReleaseAwdcShipment(v *AwdcShipment) {
	v.ArrivalCity = ""
	v.ArrivalDate = ""
	v.DepartureCity = ""
	v.DepartureDate = ""
	v.DoIn = ""
	v.DoNumber = ""
	v.DoOut = ""
	v.LogisticNumber = ""
	v.ShipmentNumber = ""
	v.Shipper = ""
	poolAwdcShipment.Put(v)
}
