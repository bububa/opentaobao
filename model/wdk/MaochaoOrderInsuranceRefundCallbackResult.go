package wdk

import (
	"sync"
)

// MaochaoOrderInsuranceRefundCallbackResult 结构体
type MaochaoOrderInsuranceRefundCallbackResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMaochaoOrderInsuranceRefundCallbackResult = sync.Pool{
	New: func() any {
		return new(MaochaoOrderInsuranceRefundCallbackResult)
	},
}

// GetMaochaoOrderInsuranceRefundCallbackResult() 从对象池中获取MaochaoOrderInsuranceRefundCallbackResult
func GetMaochaoOrderInsuranceRefundCallbackResult() *MaochaoOrderInsuranceRefundCallbackResult {
	return poolMaochaoOrderInsuranceRefundCallbackResult.Get().(*MaochaoOrderInsuranceRefundCallbackResult)
}

// ReleaseMaochaoOrderInsuranceRefundCallbackResult 释放MaochaoOrderInsuranceRefundCallbackResult
func ReleaseMaochaoOrderInsuranceRefundCallbackResult(v *MaochaoOrderInsuranceRefundCallbackResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.Success = false
	poolMaochaoOrderInsuranceRefundCallbackResult.Put(v)
}
