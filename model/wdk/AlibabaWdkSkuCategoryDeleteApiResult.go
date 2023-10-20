package wdk

import (
	"sync"
)

// AlibabaWdkSkuCategoryDeleteApiResult 结构体
type AlibabaWdkSkuCategoryDeleteApiResult struct {
	// 错误码（只有有异常才有值）
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息（只有有异常才有值）
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用成功时的返回类目code
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 接口返回成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuCategoryDeleteApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCategoryDeleteApiResult)
	},
}

// GetAlibabaWdkSkuCategoryDeleteApiResult() 从对象池中获取AlibabaWdkSkuCategoryDeleteApiResult
func GetAlibabaWdkSkuCategoryDeleteApiResult() *AlibabaWdkSkuCategoryDeleteApiResult {
	return poolAlibabaWdkSkuCategoryDeleteApiResult.Get().(*AlibabaWdkSkuCategoryDeleteApiResult)
}

// ReleaseAlibabaWdkSkuCategoryDeleteApiResult 释放AlibabaWdkSkuCategoryDeleteApiResult
func ReleaseAlibabaWdkSkuCategoryDeleteApiResult(v *AlibabaWdkSkuCategoryDeleteApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = ""
	v.Success = false
	poolAlibabaWdkSkuCategoryDeleteApiResult.Put(v)
}
