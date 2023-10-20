package wdk

import (
	"sync"
)

// AlibabaAelophyShopUpdatestatusApiResult 结构体
type AlibabaAelophyShopUpdatestatusApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAelophyShopUpdatestatusApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyShopUpdatestatusApiResult)
	},
}

// GetAlibabaAelophyShopUpdatestatusApiResult() 从对象池中获取AlibabaAelophyShopUpdatestatusApiResult
func GetAlibabaAelophyShopUpdatestatusApiResult() *AlibabaAelophyShopUpdatestatusApiResult {
	return poolAlibabaAelophyShopUpdatestatusApiResult.Get().(*AlibabaAelophyShopUpdatestatusApiResult)
}

// ReleaseAlibabaAelophyShopUpdatestatusApiResult 释放AlibabaAelophyShopUpdatestatusApiResult
func ReleaseAlibabaAelophyShopUpdatestatusApiResult(v *AlibabaAelophyShopUpdatestatusApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAelophyShopUpdatestatusApiResult.Put(v)
}
