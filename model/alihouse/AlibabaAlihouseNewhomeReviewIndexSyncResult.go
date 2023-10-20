package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeReviewIndexSyncResult 结构体
type AlibabaAlihouseNewhomeReviewIndexSyncResult struct {
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseNewhomeReviewIndexSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeReviewIndexSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeReviewIndexSyncResult() 从对象池中获取AlibabaAlihouseNewhomeReviewIndexSyncResult
func GetAlibabaAlihouseNewhomeReviewIndexSyncResult() *AlibabaAlihouseNewhomeReviewIndexSyncResult {
	return poolAlibabaAlihouseNewhomeReviewIndexSyncResult.Get().(*AlibabaAlihouseNewhomeReviewIndexSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeReviewIndexSyncResult 释放AlibabaAlihouseNewhomeReviewIndexSyncResult
func ReleaseAlibabaAlihouseNewhomeReviewIndexSyncResult(v *AlibabaAlihouseNewhomeReviewIndexSyncResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseNewhomeReviewIndexSyncResult.Put(v)
}
