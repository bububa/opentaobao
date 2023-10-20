package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomePosOpenSubmitResult 结构体
type AlibabaAlihouseExistinghomePosOpenSubmitResult struct {
	// 返回编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回值
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomePosOpenSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomePosOpenSubmitResult)
	},
}

// GetAlibabaAlihouseExistinghomePosOpenSubmitResult() 从对象池中获取AlibabaAlihouseExistinghomePosOpenSubmitResult
func GetAlibabaAlihouseExistinghomePosOpenSubmitResult() *AlibabaAlihouseExistinghomePosOpenSubmitResult {
	return poolAlibabaAlihouseExistinghomePosOpenSubmitResult.Get().(*AlibabaAlihouseExistinghomePosOpenSubmitResult)
}

// ReleaseAlibabaAlihouseExistinghomePosOpenSubmitResult 释放AlibabaAlihouseExistinghomePosOpenSubmitResult
func ReleaseAlibabaAlihouseExistinghomePosOpenSubmitResult(v *AlibabaAlihouseExistinghomePosOpenSubmitResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomePosOpenSubmitResult.Put(v)
}
