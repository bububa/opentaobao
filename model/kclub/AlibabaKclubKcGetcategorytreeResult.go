package kclub

import (
	"sync"
)

// AlibabaKclubKcGetcategorytreeResult 结构体
type AlibabaKclubKcGetcategorytreeResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 类目数据 JSONArray
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaKclubKcGetcategorytreeResult = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcGetcategorytreeResult)
	},
}

// GetAlibabaKclubKcGetcategorytreeResult() 从对象池中获取AlibabaKclubKcGetcategorytreeResult
func GetAlibabaKclubKcGetcategorytreeResult() *AlibabaKclubKcGetcategorytreeResult {
	return poolAlibabaKclubKcGetcategorytreeResult.Get().(*AlibabaKclubKcGetcategorytreeResult)
}

// ReleaseAlibabaKclubKcGetcategorytreeResult 释放AlibabaKclubKcGetcategorytreeResult
func ReleaseAlibabaKclubKcGetcategorytreeResult(v *AlibabaKclubKcGetcategorytreeResult) {
	v.Message = ""
	v.Data = ""
	v.Code = ""
	v.Success = false
	poolAlibabaKclubKcGetcategorytreeResult.Put(v)
}
