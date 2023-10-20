package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeEntrustsellingUpdateResult 结构体
type AlibabaAlihouseExistinghomeEntrustsellingUpdateResult struct {
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeEntrustsellingUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeEntrustsellingUpdateResult)
	},
}

// GetAlibabaAlihouseExistinghomeEntrustsellingUpdateResult() 从对象池中获取AlibabaAlihouseExistinghomeEntrustsellingUpdateResult
func GetAlibabaAlihouseExistinghomeEntrustsellingUpdateResult() *AlibabaAlihouseExistinghomeEntrustsellingUpdateResult {
	return poolAlibabaAlihouseExistinghomeEntrustsellingUpdateResult.Get().(*AlibabaAlihouseExistinghomeEntrustsellingUpdateResult)
}

// ReleaseAlibabaAlihouseExistinghomeEntrustsellingUpdateResult 释放AlibabaAlihouseExistinghomeEntrustsellingUpdateResult
func ReleaseAlibabaAlihouseExistinghomeEntrustsellingUpdateResult(v *AlibabaAlihouseExistinghomeEntrustsellingUpdateResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = false
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeEntrustsellingUpdateResult.Put(v)
}
