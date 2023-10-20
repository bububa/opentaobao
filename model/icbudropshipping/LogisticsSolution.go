package icbudropshipping

import (
	"sync"
)

// LogisticsSolution 结构体
type LogisticsSolution struct {
	// delivery time (days)
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// destination country
	DestinationCountry string `json:"destination_country,omitempty" xml:"destination_country,omitempty"`
	// dispatch country
	DispatchCountry string `json:"dispatch_country,omitempty" xml:"dispatch_country,omitempty"`
	// shipping type
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
	// trade term
	TradeTerm string `json:"trade_term,omitempty" xml:"trade_term,omitempty"`
	// vendor code
	VendorCode string `json:"vendor_code,omitempty" xml:"vendor_code,omitempty"`
	// vendor name
	VendorName string `json:"vendor_name,omitempty" xml:"vendor_name,omitempty"`
	// shipping fee
	Fee float64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

var poolLogisticsSolution = sync.Pool{
	New: func() any {
		return new(LogisticsSolution)
	},
}

// GetLogisticsSolution() 从对象池中获取LogisticsSolution
func GetLogisticsSolution() *LogisticsSolution {
	return poolLogisticsSolution.Get().(*LogisticsSolution)
}

// ReleaseLogisticsSolution 释放LogisticsSolution
func ReleaseLogisticsSolution(v *LogisticsSolution) {
	v.DeliveryTime = ""
	v.DestinationCountry = ""
	v.DispatchCountry = ""
	v.ShippingType = ""
	v.TradeTerm = ""
	v.VendorCode = ""
	v.VendorName = ""
	v.Fee = 0
	poolLogisticsSolution.Put(v)
}
