package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeThreedimensionSyncResult 结构体
type AlibabaAlihouseExistinghomeThreedimensionSyncResult struct {
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeThreedimensionSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeThreedimensionSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeThreedimensionSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeThreedimensionSyncResult
func GetAlibabaAlihouseExistinghomeThreedimensionSyncResult() *AlibabaAlihouseExistinghomeThreedimensionSyncResult {
	return poolAlibabaAlihouseExistinghomeThreedimensionSyncResult.Get().(*AlibabaAlihouseExistinghomeThreedimensionSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeThreedimensionSyncResult 释放AlibabaAlihouseExistinghomeThreedimensionSyncResult
func ReleaseAlibabaAlihouseExistinghomeThreedimensionSyncResult(v *AlibabaAlihouseExistinghomeThreedimensionSyncResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = false
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeThreedimensionSyncResult.Put(v)
}
