package ascpchannel

import (
	"sync"
)

// WaybillGenReceiver 结构体
type WaybillGenReceiver struct {
	// 联系人电话
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 联系人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 省
	ReceiverProvince string `json:"receiver_province,omitempty" xml:"receiver_province,omitempty"`
	// 市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 区
	ReceiverArea string `json:"receiver_area,omitempty" xml:"receiver_area,omitempty"`
}

var poolWaybillGenReceiver = sync.Pool{
	New: func() any {
		return new(WaybillGenReceiver)
	},
}

// GetWaybillGenReceiver() 从对象池中获取WaybillGenReceiver
func GetWaybillGenReceiver() *WaybillGenReceiver {
	return poolWaybillGenReceiver.Get().(*WaybillGenReceiver)
}

// ReleaseWaybillGenReceiver 释放WaybillGenReceiver
func ReleaseWaybillGenReceiver(v *WaybillGenReceiver) {
	v.ContactMobile = ""
	v.ContactName = ""
	v.DetailAddress = ""
	v.ReceiverProvince = ""
	v.ReceiverCity = ""
	v.ReceiverArea = ""
	poolWaybillGenReceiver.Put(v)
}
