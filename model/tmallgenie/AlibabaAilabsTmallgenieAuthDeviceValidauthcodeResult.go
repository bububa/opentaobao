package tmallgenie

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult struct {
	// uuid
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult() 从对象池中获取AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult
func GetAlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult() *AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult {
	return poolAlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult.Get().(*AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult 释放AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult
func ReleaseAlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult(v *AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult) {
	v.Result = ""
	v.Message = ""
	v.Code = 0
	poolAlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult.Put(v)
}
