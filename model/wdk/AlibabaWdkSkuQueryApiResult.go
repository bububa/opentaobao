package wdk

import (
	"sync"
)

// AlibabaWdkSkuQueryApiResult 结构体
type AlibabaWdkSkuQueryApiResult struct {
	// 请求参数不能为空
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 单条错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 单条查询结果
	Model *SkuDo `json:"model,omitempty" xml:"model,omitempty"`
	// 单条是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuQueryApiResult)
	},
}

// GetAlibabaWdkSkuQueryApiResult() 从对象池中获取AlibabaWdkSkuQueryApiResult
func GetAlibabaWdkSkuQueryApiResult() *AlibabaWdkSkuQueryApiResult {
	return poolAlibabaWdkSkuQueryApiResult.Get().(*AlibabaWdkSkuQueryApiResult)
}

// ReleaseAlibabaWdkSkuQueryApiResult 释放AlibabaWdkSkuQueryApiResult
func ReleaseAlibabaWdkSkuQueryApiResult(v *AlibabaWdkSkuQueryApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkSkuQueryApiResult.Put(v)
}
