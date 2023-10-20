package wdk

import (
	"sync"
)

// AlibabaWdkSeriesSortApiResult 结构体
type AlibabaWdkSeriesSortApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误详情
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 系列品重置默认商品成功结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSeriesSortApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesSortApiResult)
	},
}

// GetAlibabaWdkSeriesSortApiResult() 从对象池中获取AlibabaWdkSeriesSortApiResult
func GetAlibabaWdkSeriesSortApiResult() *AlibabaWdkSeriesSortApiResult {
	return poolAlibabaWdkSeriesSortApiResult.Get().(*AlibabaWdkSeriesSortApiResult)
}

// ReleaseAlibabaWdkSeriesSortApiResult 释放AlibabaWdkSeriesSortApiResult
func ReleaseAlibabaWdkSeriesSortApiResult(v *AlibabaWdkSeriesSortApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = false
	v.Success = false
	poolAlibabaWdkSeriesSortApiResult.Put(v)
}
