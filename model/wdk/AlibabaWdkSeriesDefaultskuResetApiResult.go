package wdk

import (
	"sync"
)

// AlibabaWdkSeriesDefaultskuResetApiResult 结构体
type AlibabaWdkSeriesDefaultskuResetApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误详情
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 系列品重置默认商品成功结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSeriesDefaultskuResetApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesDefaultskuResetApiResult)
	},
}

// GetAlibabaWdkSeriesDefaultskuResetApiResult() 从对象池中获取AlibabaWdkSeriesDefaultskuResetApiResult
func GetAlibabaWdkSeriesDefaultskuResetApiResult() *AlibabaWdkSeriesDefaultskuResetApiResult {
	return poolAlibabaWdkSeriesDefaultskuResetApiResult.Get().(*AlibabaWdkSeriesDefaultskuResetApiResult)
}

// ReleaseAlibabaWdkSeriesDefaultskuResetApiResult 释放AlibabaWdkSeriesDefaultskuResetApiResult
func ReleaseAlibabaWdkSeriesDefaultskuResetApiResult(v *AlibabaWdkSeriesDefaultskuResetApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = false
	v.Success = false
	poolAlibabaWdkSeriesDefaultskuResetApiResult.Put(v)
}
