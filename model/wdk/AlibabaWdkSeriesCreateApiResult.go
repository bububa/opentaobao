package wdk

import (
	"sync"
)

// AlibabaWdkSeriesCreateApiResult 结构体
type AlibabaWdkSeriesCreateApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误详情
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 商品系列创建结果
	Model *SkuSeriesCreateResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSeriesCreateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesCreateApiResult)
	},
}

// GetAlibabaWdkSeriesCreateApiResult() 从对象池中获取AlibabaWdkSeriesCreateApiResult
func GetAlibabaWdkSeriesCreateApiResult() *AlibabaWdkSeriesCreateApiResult {
	return poolAlibabaWdkSeriesCreateApiResult.Get().(*AlibabaWdkSeriesCreateApiResult)
}

// ReleaseAlibabaWdkSeriesCreateApiResult 释放AlibabaWdkSeriesCreateApiResult
func ReleaseAlibabaWdkSeriesCreateApiResult(v *AlibabaWdkSeriesCreateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkSeriesCreateApiResult.Put(v)
}
