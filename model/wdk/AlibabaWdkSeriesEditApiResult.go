package wdk

import (
	"sync"
)

// AlibabaWdkSeriesEditApiResult 结构体
type AlibabaWdkSeriesEditApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误详情
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 商品系列更新成功结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSeriesEditApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesEditApiResult)
	},
}

// GetAlibabaWdkSeriesEditApiResult() 从对象池中获取AlibabaWdkSeriesEditApiResult
func GetAlibabaWdkSeriesEditApiResult() *AlibabaWdkSeriesEditApiResult {
	return poolAlibabaWdkSeriesEditApiResult.Get().(*AlibabaWdkSeriesEditApiResult)
}

// ReleaseAlibabaWdkSeriesEditApiResult 释放AlibabaWdkSeriesEditApiResult
func ReleaseAlibabaWdkSeriesEditApiResult(v *AlibabaWdkSeriesEditApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = false
	v.Success = false
	poolAlibabaWdkSeriesEditApiResult.Put(v)
}
