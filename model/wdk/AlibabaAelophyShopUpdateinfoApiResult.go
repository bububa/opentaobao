package wdk

import (
	"sync"
)

// AlibabaAelophyShopUpdateinfoApiResult 结构体
type AlibabaAelophyShopUpdateinfoApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAelophyShopUpdateinfoApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyShopUpdateinfoApiResult)
	},
}

// GetAlibabaAelophyShopUpdateinfoApiResult() 从对象池中获取AlibabaAelophyShopUpdateinfoApiResult
func GetAlibabaAelophyShopUpdateinfoApiResult() *AlibabaAelophyShopUpdateinfoApiResult {
	return poolAlibabaAelophyShopUpdateinfoApiResult.Get().(*AlibabaAelophyShopUpdateinfoApiResult)
}

// ReleaseAlibabaAelophyShopUpdateinfoApiResult 释放AlibabaAelophyShopUpdateinfoApiResult
func ReleaseAlibabaAelophyShopUpdateinfoApiResult(v *AlibabaAelophyShopUpdateinfoApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAelophyShopUpdateinfoApiResult.Put(v)
}
