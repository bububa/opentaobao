package tmallgenie

import (
	"sync"
)

// DeviceSecretInfo 结构体
type DeviceSecretInfo struct {
	// 错误原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 秘钥
	Secret string `json:"secret,omitempty" xml:"secret,omitempty"`
	// mac地址
	Mac string `json:"mac,omitempty" xml:"mac,omitempty"`
	// md5
	Md5 string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 获取秘钥是否成功
	Sucess bool `json:"sucess,omitempty" xml:"sucess,omitempty"`
}

var poolDeviceSecretInfo = sync.Pool{
	New: func() any {
		return new(DeviceSecretInfo)
	},
}

// GetDeviceSecretInfo() 从对象池中获取DeviceSecretInfo
func GetDeviceSecretInfo() *DeviceSecretInfo {
	return poolDeviceSecretInfo.Get().(*DeviceSecretInfo)
}

// ReleaseDeviceSecretInfo 释放DeviceSecretInfo
func ReleaseDeviceSecretInfo(v *DeviceSecretInfo) {
	v.Reason = ""
	v.Secret = ""
	v.Mac = ""
	v.Md5 = ""
	v.Sucess = false
	poolDeviceSecretInfo.Put(v)
}
