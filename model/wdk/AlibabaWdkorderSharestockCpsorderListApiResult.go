package wdk

import (
	"sync"
)

// AlibabaWdkorderSharestockCpsorderListApiResult 结构体
type AlibabaWdkorderSharestockCpsorderListApiResult struct {
	// 调用接口返回对象
	Model []CpsOrderResponse `json:"model,omitempty" xml:"model>cps_order_response,omitempty"`
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口返回错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用接口成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkorderSharestockCpsorderListApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkorderSharestockCpsorderListApiResult)
	},
}

// GetAlibabaWdkorderSharestockCpsorderListApiResult() 从对象池中获取AlibabaWdkorderSharestockCpsorderListApiResult
func GetAlibabaWdkorderSharestockCpsorderListApiResult() *AlibabaWdkorderSharestockCpsorderListApiResult {
	return poolAlibabaWdkorderSharestockCpsorderListApiResult.Get().(*AlibabaWdkorderSharestockCpsorderListApiResult)
}

// ReleaseAlibabaWdkorderSharestockCpsorderListApiResult 释放AlibabaWdkorderSharestockCpsorderListApiResult
func ReleaseAlibabaWdkorderSharestockCpsorderListApiResult(v *AlibabaWdkorderSharestockCpsorderListApiResult) {
	v.Model = v.Model[:0]
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaWdkorderSharestockCpsorderListApiResult.Put(v)
}
