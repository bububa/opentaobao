package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeReviewSyncResult 结构体
type AlibabaAlihouseNewhomeReviewSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseNewhomeReviewSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeReviewSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeReviewSyncResult() 从对象池中获取AlibabaAlihouseNewhomeReviewSyncResult
func GetAlibabaAlihouseNewhomeReviewSyncResult() *AlibabaAlihouseNewhomeReviewSyncResult {
	return poolAlibabaAlihouseNewhomeReviewSyncResult.Get().(*AlibabaAlihouseNewhomeReviewSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeReviewSyncResult 释放AlibabaAlihouseNewhomeReviewSyncResult
func ReleaseAlibabaAlihouseNewhomeReviewSyncResult(v *AlibabaAlihouseNewhomeReviewSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseNewhomeReviewSyncResult.Put(v)
}
