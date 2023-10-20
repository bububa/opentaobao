package wdk

import (
	"sync"
)

// AlibabaAelophyShopUpdaterangeApiResult 结构体
type AlibabaAelophyShopUpdaterangeApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAelophyShopUpdaterangeApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyShopUpdaterangeApiResult)
	},
}

// GetAlibabaAelophyShopUpdaterangeApiResult() 从对象池中获取AlibabaAelophyShopUpdaterangeApiResult
func GetAlibabaAelophyShopUpdaterangeApiResult() *AlibabaAelophyShopUpdaterangeApiResult {
	return poolAlibabaAelophyShopUpdaterangeApiResult.Get().(*AlibabaAelophyShopUpdaterangeApiResult)
}

// ReleaseAlibabaAelophyShopUpdaterangeApiResult 释放AlibabaAelophyShopUpdaterangeApiResult
func ReleaseAlibabaAelophyShopUpdaterangeApiResult(v *AlibabaAelophyShopUpdaterangeApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAelophyShopUpdaterangeApiResult.Put(v)
}
