package alilabs

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult struct {
	// 结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 接口返回model
	Result *DeviceStatusVo `json:"result,omitempty" xml:"result,omitempty"`
	// 状态码（200：成功，其他：失败）
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult() 从对象池中获取AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult
func GetAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult() *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult {
	return poolAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult.Get().(*AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult 释放AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult
func ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult(v *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult) {
	v.Message = ""
	v.Result = nil
	v.Code = 0
	poolAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult.Put(v)
}
