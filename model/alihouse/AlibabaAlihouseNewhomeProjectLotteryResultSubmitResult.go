package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult 结构体
type AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult struct {
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回消息体
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectLotteryResultSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult
func GetAlibabaAlihouseNewhomeProjectLotteryResultSubmitResult() *AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult {
	return poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitResult.Get().(*AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectLotteryResultSubmitResult 释放AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult
func ReleaseAlibabaAlihouseNewhomeProjectLotteryResultSubmitResult(v *AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitResult.Put(v)
}
