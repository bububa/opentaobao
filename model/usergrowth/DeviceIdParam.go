package usergrowth

import (
	"sync"
)

// DeviceIdParam 结构体
type DeviceIdParam struct {
	// 手机系统
	Os string `json:"os,omitempty" xml:"os,omitempty"`
	// 设备类型 oaid/caid/imei
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 设备值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 是否是md5值
	IsMd5 bool `json:"is_md5,omitempty" xml:"is_md5,omitempty"`
}

var poolDeviceIdParam = sync.Pool{
	New: func() any {
		return new(DeviceIdParam)
	},
}

// GetDeviceIdParam() 从对象池中获取DeviceIdParam
func GetDeviceIdParam() *DeviceIdParam {
	return poolDeviceIdParam.Get().(*DeviceIdParam)
}

// ReleaseDeviceIdParam 释放DeviceIdParam
func ReleaseDeviceIdParam(v *DeviceIdParam) {
	v.Os = ""
	v.Type = ""
	v.Value = ""
	v.IsMd5 = false
	poolDeviceIdParam.Put(v)
}
