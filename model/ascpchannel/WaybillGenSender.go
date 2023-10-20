package ascpchannel

import (
	"sync"
)

// WaybillGenSender 结构体
type WaybillGenSender struct {
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 省
	SenderProvince string `json:"sender_province,omitempty" xml:"sender_province,omitempty"`
	// 市
	SenderCity string `json:"sender_city,omitempty" xml:"sender_city,omitempty"`
	// 区
	SenderArea string `json:"sender_area,omitempty" xml:"sender_area,omitempty"`
}

var poolWaybillGenSender = sync.Pool{
	New: func() any {
		return new(WaybillGenSender)
	},
}

// GetWaybillGenSender() 从对象池中获取WaybillGenSender
func GetWaybillGenSender() *WaybillGenSender {
	return poolWaybillGenSender.Get().(*WaybillGenSender)
}

// ReleaseWaybillGenSender 释放WaybillGenSender
func ReleaseWaybillGenSender(v *WaybillGenSender) {
	v.DetailAddress = ""
	v.SenderProvince = ""
	v.SenderCity = ""
	v.SenderArea = ""
	poolWaybillGenSender.Put(v)
}
