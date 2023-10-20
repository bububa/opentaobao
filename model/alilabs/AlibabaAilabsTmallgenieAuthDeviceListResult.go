package alilabs

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthDeviceListResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceListResult struct {
	// 拓展信息
	Extensions string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// 设备名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 设备唯一id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 设备UUID
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthDeviceListResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceListResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceListResult() 从对象池中获取AlibabaAilabsTmallgenieAuthDeviceListResult
func GetAlibabaAilabsTmallgenieAuthDeviceListResult() *AlibabaAilabsTmallgenieAuthDeviceListResult {
	return poolAlibabaAilabsTmallgenieAuthDeviceListResult.Get().(*AlibabaAilabsTmallgenieAuthDeviceListResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceListResult 释放AlibabaAilabsTmallgenieAuthDeviceListResult
func ReleaseAlibabaAilabsTmallgenieAuthDeviceListResult(v *AlibabaAilabsTmallgenieAuthDeviceListResult) {
	v.Extensions = ""
	v.Name = ""
	v.DeviceId = ""
	v.Uuid = ""
	poolAlibabaAilabsTmallgenieAuthDeviceListResult.Put(v)
}
