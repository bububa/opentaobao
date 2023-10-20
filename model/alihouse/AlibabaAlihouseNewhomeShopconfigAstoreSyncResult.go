package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeShopconfigAstoreSyncResult 结构体
type AlibabaAlihouseNewhomeShopconfigAstoreSyncResult struct {
	// 处理完成后的消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误/成功码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回的结果
	Data *AstoreRespDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeShopconfigAstoreSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeShopconfigAstoreSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeShopconfigAstoreSyncResult() 从对象池中获取AlibabaAlihouseNewhomeShopconfigAstoreSyncResult
func GetAlibabaAlihouseNewhomeShopconfigAstoreSyncResult() *AlibabaAlihouseNewhomeShopconfigAstoreSyncResult {
	return poolAlibabaAlihouseNewhomeShopconfigAstoreSyncResult.Get().(*AlibabaAlihouseNewhomeShopconfigAstoreSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigAstoreSyncResult 释放AlibabaAlihouseNewhomeShopconfigAstoreSyncResult
func ReleaseAlibabaAlihouseNewhomeShopconfigAstoreSyncResult(v *AlibabaAlihouseNewhomeShopconfigAstoreSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAlihouseNewhomeShopconfigAstoreSyncResult.Put(v)
}
