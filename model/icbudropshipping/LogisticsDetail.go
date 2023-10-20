package icbudropshipping

import (
	"sync"
)

// LogisticsDetail 结构体
type LogisticsDetail struct {
	// For BuyNow orders, use the value shown in the vendorCode field from the shipping cost template API; non-BuyNow orders don’t need to provide this information.    alibaba.shipping.freight.calculate  &#39;s vender_code
	CarrierCode string `json:"carrier_code,omitempty" xml:"carrier_code,omitempty"`
	// shipment address
	ShipmentAddress *Address `json:"shipment_address,omitempty" xml:"shipment_address,omitempty"`
}

var poolLogisticsDetail = sync.Pool{
	New: func() any {
		return new(LogisticsDetail)
	},
}

// GetLogisticsDetail() 从对象池中获取LogisticsDetail
func GetLogisticsDetail() *LogisticsDetail {
	return poolLogisticsDetail.Get().(*LogisticsDetail)
}

// ReleaseLogisticsDetail 释放LogisticsDetail
func ReleaseLogisticsDetail(v *LogisticsDetail) {
	v.CarrierCode = ""
	v.ShipmentAddress = nil
	poolLogisticsDetail.Put(v)
}
