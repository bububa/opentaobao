package alilabs

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthDeviceGetResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceGetResult struct {
	// 拓展信息
	Extensions string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// 设备名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 设备三方唯一id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 设备uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 设备标签
	DeviceTag *HashMap `json:"device_tag,omitempty" xml:"device_tag,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthDeviceGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceGetResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceGetResult() 从对象池中获取AlibabaAilabsTmallgenieAuthDeviceGetResult
func GetAlibabaAilabsTmallgenieAuthDeviceGetResult() *AlibabaAilabsTmallgenieAuthDeviceGetResult {
	return poolAlibabaAilabsTmallgenieAuthDeviceGetResult.Get().(*AlibabaAilabsTmallgenieAuthDeviceGetResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceGetResult 释放AlibabaAilabsTmallgenieAuthDeviceGetResult
func ReleaseAlibabaAilabsTmallgenieAuthDeviceGetResult(v *AlibabaAilabsTmallgenieAuthDeviceGetResult) {
	v.Extensions = ""
	v.Name = ""
	v.DeviceId = ""
	v.Uuid = ""
	v.DeviceTag = nil
	poolAlibabaAilabsTmallgenieAuthDeviceGetResult.Put(v)
}
