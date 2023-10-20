package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeReviewChangestatusResult 结构体
type AlibabaAlihouseNewhomeReviewChangestatusResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 更新结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeReviewChangestatusResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeReviewChangestatusResult)
	},
}

// GetAlibabaAlihouseNewhomeReviewChangestatusResult() 从对象池中获取AlibabaAlihouseNewhomeReviewChangestatusResult
func GetAlibabaAlihouseNewhomeReviewChangestatusResult() *AlibabaAlihouseNewhomeReviewChangestatusResult {
	return poolAlibabaAlihouseNewhomeReviewChangestatusResult.Get().(*AlibabaAlihouseNewhomeReviewChangestatusResult)
}

// ReleaseAlibabaAlihouseNewhomeReviewChangestatusResult 释放AlibabaAlihouseNewhomeReviewChangestatusResult
func ReleaseAlibabaAlihouseNewhomeReviewChangestatusResult(v *AlibabaAlihouseNewhomeReviewChangestatusResult) {
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseNewhomeReviewChangestatusResult.Put(v)
}
