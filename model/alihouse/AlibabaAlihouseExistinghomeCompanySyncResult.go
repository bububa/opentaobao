package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeCompanySyncResult 结构体
type AlibabaAlihouseExistinghomeCompanySyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeCompanySyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeCompanySyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeCompanySyncResult() 从对象池中获取AlibabaAlihouseExistinghomeCompanySyncResult
func GetAlibabaAlihouseExistinghomeCompanySyncResult() *AlibabaAlihouseExistinghomeCompanySyncResult {
	return poolAlibabaAlihouseExistinghomeCompanySyncResult.Get().(*AlibabaAlihouseExistinghomeCompanySyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeCompanySyncResult 释放AlibabaAlihouseExistinghomeCompanySyncResult
func ReleaseAlibabaAlihouseExistinghomeCompanySyncResult(v *AlibabaAlihouseExistinghomeCompanySyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseExistinghomeCompanySyncResult.Put(v)
}
