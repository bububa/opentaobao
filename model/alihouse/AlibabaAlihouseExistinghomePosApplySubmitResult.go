package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomePosApplySubmitResult 结构体
type AlibabaAlihouseExistinghomePosApplySubmitResult struct {
	// 编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回值
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomePosApplySubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomePosApplySubmitResult)
	},
}

// GetAlibabaAlihouseExistinghomePosApplySubmitResult() 从对象池中获取AlibabaAlihouseExistinghomePosApplySubmitResult
func GetAlibabaAlihouseExistinghomePosApplySubmitResult() *AlibabaAlihouseExistinghomePosApplySubmitResult {
	return poolAlibabaAlihouseExistinghomePosApplySubmitResult.Get().(*AlibabaAlihouseExistinghomePosApplySubmitResult)
}

// ReleaseAlibabaAlihouseExistinghomePosApplySubmitResult 释放AlibabaAlihouseExistinghomePosApplySubmitResult
func ReleaseAlibabaAlihouseExistinghomePosApplySubmitResult(v *AlibabaAlihouseExistinghomePosApplySubmitResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomePosApplySubmitResult.Put(v)
}
