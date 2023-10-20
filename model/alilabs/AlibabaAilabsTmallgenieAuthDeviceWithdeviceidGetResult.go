package alilabs

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 用户安全ID
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
	// 设备名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 天猫精灵设备唯一ID
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 设备Mac地址
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 结果
	Result *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult `json:"result,omitempty" xml:"result,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult() 从对象池中获取AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult
func GetAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult() *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult {
	return poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult.Get().(*AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult 释放AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult
func ReleaseAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult(v *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult) {
	v.Message = ""
	v.UserOpenId = ""
	v.Name = ""
	v.Uuid = ""
	v.DeviceId = ""
	v.Result = nil
	v.Code = 0
	poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult.Put(v)
}
