package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeNewReviewSyncResult 结构体
type AlibabaAlihouseNewhomeNewReviewSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseNewhomeNewReviewSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeNewReviewSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeNewReviewSyncResult() 从对象池中获取AlibabaAlihouseNewhomeNewReviewSyncResult
func GetAlibabaAlihouseNewhomeNewReviewSyncResult() *AlibabaAlihouseNewhomeNewReviewSyncResult {
	return poolAlibabaAlihouseNewhomeNewReviewSyncResult.Get().(*AlibabaAlihouseNewhomeNewReviewSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeNewReviewSyncResult 释放AlibabaAlihouseNewhomeNewReviewSyncResult
func ReleaseAlibabaAlihouseNewhomeNewReviewSyncResult(v *AlibabaAlihouseNewhomeNewReviewSyncResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseNewhomeNewReviewSyncResult.Put(v)
}
