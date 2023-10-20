package wdk

import (
	"sync"
)

// MaochaoOrderInsuranceCallbackResult 结构体
type MaochaoOrderInsuranceCallbackResult struct {
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMaochaoOrderInsuranceCallbackResult = sync.Pool{
	New: func() any {
		return new(MaochaoOrderInsuranceCallbackResult)
	},
}

// GetMaochaoOrderInsuranceCallbackResult() 从对象池中获取MaochaoOrderInsuranceCallbackResult
func GetMaochaoOrderInsuranceCallbackResult() *MaochaoOrderInsuranceCallbackResult {
	return poolMaochaoOrderInsuranceCallbackResult.Get().(*MaochaoOrderInsuranceCallbackResult)
}

// ReleaseMaochaoOrderInsuranceCallbackResult 释放MaochaoOrderInsuranceCallbackResult
func ReleaseMaochaoOrderInsuranceCallbackResult(v *MaochaoOrderInsuranceCallbackResult) {
	v.ReturnMsg = ""
	v.ReturnCode = ""
	v.Success = false
	poolMaochaoOrderInsuranceCallbackResult.Put(v)
}
