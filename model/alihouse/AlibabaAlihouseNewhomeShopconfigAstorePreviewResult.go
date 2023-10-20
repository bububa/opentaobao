package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeShopconfigAstorePreviewResult 结构体
type AlibabaAlihouseNewhomeShopconfigAstorePreviewResult struct {
	// 返回的结果
	Data []AstoreRespDto `json:"data,omitempty" xml:"data>astore_resp_dto,omitempty"`
	// 处理完成后的消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误/成功码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeShopconfigAstorePreviewResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeShopconfigAstorePreviewResult)
	},
}

// GetAlibabaAlihouseNewhomeShopconfigAstorePreviewResult() 从对象池中获取AlibabaAlihouseNewhomeShopconfigAstorePreviewResult
func GetAlibabaAlihouseNewhomeShopconfigAstorePreviewResult() *AlibabaAlihouseNewhomeShopconfigAstorePreviewResult {
	return poolAlibabaAlihouseNewhomeShopconfigAstorePreviewResult.Get().(*AlibabaAlihouseNewhomeShopconfigAstorePreviewResult)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigAstorePreviewResult 释放AlibabaAlihouseNewhomeShopconfigAstorePreviewResult
func ReleaseAlibabaAlihouseNewhomeShopconfigAstorePreviewResult(v *AlibabaAlihouseNewhomeShopconfigAstorePreviewResult) {
	v.Data = v.Data[:0]
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeShopconfigAstorePreviewResult.Put(v)
}
