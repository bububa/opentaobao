package wdk

import (
	"sync"
)

// AlibabaWdkSeriesSkuAddApiResult 结构体
type AlibabaWdkSeriesSkuAddApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误详情
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 系列品添加商品成功结果
	Model *SkuSeriesEditResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSeriesSkuAddApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSeriesSkuAddApiResult)
	},
}

// GetAlibabaWdkSeriesSkuAddApiResult() 从对象池中获取AlibabaWdkSeriesSkuAddApiResult
func GetAlibabaWdkSeriesSkuAddApiResult() *AlibabaWdkSeriesSkuAddApiResult {
	return poolAlibabaWdkSeriesSkuAddApiResult.Get().(*AlibabaWdkSeriesSkuAddApiResult)
}

// ReleaseAlibabaWdkSeriesSkuAddApiResult 释放AlibabaWdkSeriesSkuAddApiResult
func ReleaseAlibabaWdkSeriesSkuAddApiResult(v *AlibabaWdkSeriesSkuAddApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkSeriesSkuAddApiResult.Put(v)
}
