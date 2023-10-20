package xhotelonlineorder

import (
	"sync"
)

// FaceScanParam 结构体
type FaceScanParam struct {
	// 姓名
	GuestName string `json:"guest_name,omitempty" xml:"guest_name,omitempty"`
	// 证件号
	CertificateNum string `json:"certificate_num,omitempty" xml:"certificate_num,omitempty"`
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 设备厂商名，下划线隔开，字母小写
	DeviceFirm string `json:"device_firm,omitempty" xml:"device_firm,omitempty"`
	// 证件类型：20：身份证
	CertificateType int64 `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
}

var poolFaceScanParam = sync.Pool{
	New: func() any {
		return new(FaceScanParam)
	},
}

// GetFaceScanParam() 从对象池中获取FaceScanParam
func GetFaceScanParam() *FaceScanParam {
	return poolFaceScanParam.Get().(*FaceScanParam)
}

// ReleaseFaceScanParam 释放FaceScanParam
func ReleaseFaceScanParam(v *FaceScanParam) {
	v.GuestName = ""
	v.CertificateNum = ""
	v.DeviceId = ""
	v.DeviceFirm = ""
	v.CertificateType = 0
	poolFaceScanParam.Put(v)
}
