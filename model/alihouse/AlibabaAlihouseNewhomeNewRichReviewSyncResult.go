package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeNewRichReviewSyncResult 结构体
type AlibabaAlihouseNewhomeNewRichReviewSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 123
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseNewhomeNewRichReviewSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeNewRichReviewSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeNewRichReviewSyncResult() 从对象池中获取AlibabaAlihouseNewhomeNewRichReviewSyncResult
func GetAlibabaAlihouseNewhomeNewRichReviewSyncResult() *AlibabaAlihouseNewhomeNewRichReviewSyncResult {
	return poolAlibabaAlihouseNewhomeNewRichReviewSyncResult.Get().(*AlibabaAlihouseNewhomeNewRichReviewSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeNewRichReviewSyncResult 释放AlibabaAlihouseNewhomeNewRichReviewSyncResult
func ReleaseAlibabaAlihouseNewhomeNewRichReviewSyncResult(v *AlibabaAlihouseNewhomeNewRichReviewSyncResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseNewhomeNewRichReviewSyncResult.Put(v)
}
