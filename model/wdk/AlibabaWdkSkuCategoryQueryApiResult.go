package wdk

import (
	"sync"
)

// AlibabaWdkSkuCategoryQueryApiResult 结构体
type AlibabaWdkSkuCategoryQueryApiResult struct {
	// 错误码（只有有异常才有值）
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息（只有有异常才有值）
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回业务参数模型，-1状态的为删除的类目
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 接口返回成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuCategoryQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCategoryQueryApiResult)
	},
}

// GetAlibabaWdkSkuCategoryQueryApiResult() 从对象池中获取AlibabaWdkSkuCategoryQueryApiResult
func GetAlibabaWdkSkuCategoryQueryApiResult() *AlibabaWdkSkuCategoryQueryApiResult {
	return poolAlibabaWdkSkuCategoryQueryApiResult.Get().(*AlibabaWdkSkuCategoryQueryApiResult)
}

// ReleaseAlibabaWdkSkuCategoryQueryApiResult 释放AlibabaWdkSkuCategoryQueryApiResult
func ReleaseAlibabaWdkSkuCategoryQueryApiResult(v *AlibabaWdkSkuCategoryQueryApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = ""
	v.Success = false
	poolAlibabaWdkSkuCategoryQueryApiResult.Put(v)
}
