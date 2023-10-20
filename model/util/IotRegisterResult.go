package util

import (
	"sync"
)

// IotRegisterResult 结构体
type IotRegisterResult struct {
	// 产品Key
	ProductKey string `json:"product_key,omitempty" xml:"product_key,omitempty"`
	// 设备名称
	DeviceName string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// 设备Secret
	DeviceSecret string `json:"device_secret,omitempty" xml:"device_secret,omitempty"`
}

var poolIotRegisterResult = sync.Pool{
	New: func() any {
		return new(IotRegisterResult)
	},
}

// GetIotRegisterResult() 从对象池中获取IotRegisterResult
func GetIotRegisterResult() *IotRegisterResult {
	return poolIotRegisterResult.Get().(*IotRegisterResult)
}

// ReleaseIotRegisterResult 释放IotRegisterResult
func ReleaseIotRegisterResult(v *IotRegisterResult) {
	v.ProductKey = ""
	v.DeviceName = ""
	v.DeviceSecret = ""
	poolIotRegisterResult.Put(v)
}
