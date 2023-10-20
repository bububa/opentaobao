package tmallgenie

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应结果
	Result *AuthResultVo `json:"result,omitempty" xml:"result,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult() 从对象池中获取AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult
func GetAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult() *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult {
	return poolAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult.Get().(*AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult 释放AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult
func ReleaseAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult(v *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult) {
	v.Message = ""
	v.Result = nil
	v.Code = 0
	poolAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult.Put(v)
}
