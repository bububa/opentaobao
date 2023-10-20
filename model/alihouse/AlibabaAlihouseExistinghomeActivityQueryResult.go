package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeActivityQueryResult 结构体
type AlibabaAlihouseExistinghomeActivityQueryResult struct {
	// 结果信息
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 处理完成后的消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误/成功码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeActivityQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeActivityQueryResult)
	},
}

// GetAlibabaAlihouseExistinghomeActivityQueryResult() 从对象池中获取AlibabaAlihouseExistinghomeActivityQueryResult
func GetAlibabaAlihouseExistinghomeActivityQueryResult() *AlibabaAlihouseExistinghomeActivityQueryResult {
	return poolAlibabaAlihouseExistinghomeActivityQueryResult.Get().(*AlibabaAlihouseExistinghomeActivityQueryResult)
}

// ReleaseAlibabaAlihouseExistinghomeActivityQueryResult 释放AlibabaAlihouseExistinghomeActivityQueryResult
func ReleaseAlibabaAlihouseExistinghomeActivityQueryResult(v *AlibabaAlihouseExistinghomeActivityQueryResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseExistinghomeActivityQueryResult.Put(v)
}
