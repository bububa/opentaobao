package icbudropshipping

import (
	"sync"
)

// Value 结构体
type Value struct {
	// shipping type
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
	// trade term
	TradeTerm string `json:"trade_term,omitempty" xml:"trade_term,omitempty"`
	// dispatch country
	DispatchCountry string `json:"dispatch_country,omitempty" xml:"dispatch_country,omitempty"`
	// destination country
	DestinationCountry string `json:"destination_country,omitempty" xml:"destination_country,omitempty"`
	// vendor code
	VendorCode string `json:"vendor_code,omitempty" xml:"vendor_code,omitempty"`
	// vendor name
	VendorName string `json:"vendor_name,omitempty" xml:"vendor_name,omitempty"`
	// delivery time (days)
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// shipping fee
	Fee float64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

var poolValue = sync.Pool{
	New: func() any {
		return new(Value)
	},
}

// GetValue() 从对象池中获取Value
func GetValue() *Value {
	return poolValue.Get().(*Value)
}

// ReleaseValue 释放Value
func ReleaseValue(v *Value) {
	v.ShippingType = ""
	v.TradeTerm = ""
	v.DispatchCountry = ""
	v.DestinationCountry = ""
	v.VendorCode = ""
	v.VendorName = ""
	v.DeliveryTime = ""
	v.Fee = 0
	poolValue.Put(v)
}
