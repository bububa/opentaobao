package wdk

import (
	"sync"
)

// DeliveryInfo 结构体
type DeliveryInfo struct {
	// 送货人名称
	DeliveryName string `json:"delivery_name,omitempty" xml:"delivery_name,omitempty"`
	// 送货人手机号
	DeliveryPhone string `json:"delivery_phone,omitempty" xml:"delivery_phone,omitempty"`
}

var poolDeliveryInfo = sync.Pool{
	New: func() any {
		return new(DeliveryInfo)
	},
}

// GetDeliveryInfo() 从对象池中获取DeliveryInfo
func GetDeliveryInfo() *DeliveryInfo {
	return poolDeliveryInfo.Get().(*DeliveryInfo)
}

// ReleaseDeliveryInfo 释放DeliveryInfo
func ReleaseDeliveryInfo(v *DeliveryInfo) {
	v.DeliveryName = ""
	v.DeliveryPhone = ""
	poolDeliveryInfo.Put(v)
}
