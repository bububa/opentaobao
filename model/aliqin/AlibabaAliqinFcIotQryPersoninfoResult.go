package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotQryPersoninfoResult 结构体
type AlibabaAliqinFcIotQryPersoninfoResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotQryPersoninfoResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotQryPersoninfoResult)
	},
}

// GetAlibabaAliqinFcIotQryPersoninfoResult() 从对象池中获取AlibabaAliqinFcIotQryPersoninfoResult
func GetAlibabaAliqinFcIotQryPersoninfoResult() *AlibabaAliqinFcIotQryPersoninfoResult {
	return poolAlibabaAliqinFcIotQryPersoninfoResult.Get().(*AlibabaAliqinFcIotQryPersoninfoResult)
}

// ReleaseAlibabaAliqinFcIotQryPersoninfoResult 释放AlibabaAliqinFcIotQryPersoninfoResult
func ReleaseAlibabaAliqinFcIotQryPersoninfoResult(v *AlibabaAliqinFcIotQryPersoninfoResult) {
	v.Model = ""
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinFcIotQryPersoninfoResult.Put(v)
}
