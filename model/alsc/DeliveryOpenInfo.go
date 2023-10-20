package alsc

import (
	"sync"
)

// DeliveryOpenInfo 结构体
type DeliveryOpenInfo struct {
	// 配送员名称
	DeliverMan string `json:"deliver_man,omitempty" xml:"deliver_man,omitempty"`
	// 开始配送时间
	DeliverTime string `json:"deliver_time,omitempty" xml:"deliver_time,omitempty"`
	// 物流单扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 配送员联系电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 物流费
	Fee int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

var poolDeliveryOpenInfo = sync.Pool{
	New: func() any {
		return new(DeliveryOpenInfo)
	},
}

// GetDeliveryOpenInfo() 从对象池中获取DeliveryOpenInfo
func GetDeliveryOpenInfo() *DeliveryOpenInfo {
	return poolDeliveryOpenInfo.Get().(*DeliveryOpenInfo)
}

// ReleaseDeliveryOpenInfo 释放DeliveryOpenInfo
func ReleaseDeliveryOpenInfo(v *DeliveryOpenInfo) {
	v.DeliverMan = ""
	v.DeliverTime = ""
	v.ExtInfo = ""
	v.Phone = ""
	v.Fee = 0
	poolDeliveryOpenInfo.Put(v)
}
