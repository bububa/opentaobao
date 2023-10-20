package wdk

import (
	"sync"
)

// AlibabaWdkUmsInventoryPublishApiResult 结构体
type AlibabaWdkUmsInventoryPublishApiResult struct {
	// 调用服务返回结果对象
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 调用服务返回错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用服务返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用服务返回成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkUmsInventoryPublishApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkUmsInventoryPublishApiResult)
	},
}

// GetAlibabaWdkUmsInventoryPublishApiResult() 从对象池中获取AlibabaWdkUmsInventoryPublishApiResult
func GetAlibabaWdkUmsInventoryPublishApiResult() *AlibabaWdkUmsInventoryPublishApiResult {
	return poolAlibabaWdkUmsInventoryPublishApiResult.Get().(*AlibabaWdkUmsInventoryPublishApiResult)
}

// ReleaseAlibabaWdkUmsInventoryPublishApiResult 释放AlibabaWdkUmsInventoryPublishApiResult
func ReleaseAlibabaWdkUmsInventoryPublishApiResult(v *AlibabaWdkUmsInventoryPublishApiResult) {
	v.Model = ""
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkUmsInventoryPublishApiResult.Put(v)
}
