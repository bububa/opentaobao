package logistic

import (
	"sync"
)

// CainiaoCbossWorkplatformBiztypeQuerybyidResult 结构体
type CainiaoCbossWorkplatformBiztypeQuerybyidResult struct {
	// bizTypeJson
	BizTypeJson string `json:"biz_type_json,omitempty" xml:"biz_type_json,omitempty"`
	// errorCode
	ResErrorCode string `json:"res_error_code,omitempty" xml:"res_error_code,omitempty"`
	// errorMsg
	ResErrorMsg string `json:"res_error_msg,omitempty" xml:"res_error_msg,omitempty"`
	// success
	ResSuccess bool `json:"res_success,omitempty" xml:"res_success,omitempty"`
}

var poolCainiaoCbossWorkplatformBiztypeQuerybyidResult = sync.Pool{
	New: func() any {
		return new(CainiaoCbossWorkplatformBiztypeQuerybyidResult)
	},
}

// GetCainiaoCbossWorkplatformBiztypeQuerybyidResult() 从对象池中获取CainiaoCbossWorkplatformBiztypeQuerybyidResult
func GetCainiaoCbossWorkplatformBiztypeQuerybyidResult() *CainiaoCbossWorkplatformBiztypeQuerybyidResult {
	return poolCainiaoCbossWorkplatformBiztypeQuerybyidResult.Get().(*CainiaoCbossWorkplatformBiztypeQuerybyidResult)
}

// ReleaseCainiaoCbossWorkplatformBiztypeQuerybyidResult 释放CainiaoCbossWorkplatformBiztypeQuerybyidResult
func ReleaseCainiaoCbossWorkplatformBiztypeQuerybyidResult(v *CainiaoCbossWorkplatformBiztypeQuerybyidResult) {
	v.BizTypeJson = ""
	v.ResErrorCode = ""
	v.ResErrorMsg = ""
	v.ResSuccess = false
	poolCainiaoCbossWorkplatformBiztypeQuerybyidResult.Put(v)
}
