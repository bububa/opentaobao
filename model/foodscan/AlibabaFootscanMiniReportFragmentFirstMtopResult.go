package foodscan

import (
	"sync"
)

// AlibabaFootscanMiniReportFragmentFirstMtopResult 结构体
type AlibabaFootscanMiniReportFragmentFirstMtopResult struct {
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回i数据
	Data *AlibabaFootscanMiniReportFragmentFirstData `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaFootscanMiniReportFragmentFirstMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniReportFragmentFirstMtopResult)
	},
}

// GetAlibabaFootscanMiniReportFragmentFirstMtopResult() 从对象池中获取AlibabaFootscanMiniReportFragmentFirstMtopResult
func GetAlibabaFootscanMiniReportFragmentFirstMtopResult() *AlibabaFootscanMiniReportFragmentFirstMtopResult {
	return poolAlibabaFootscanMiniReportFragmentFirstMtopResult.Get().(*AlibabaFootscanMiniReportFragmentFirstMtopResult)
}

// ReleaseAlibabaFootscanMiniReportFragmentFirstMtopResult 释放AlibabaFootscanMiniReportFragmentFirstMtopResult
func ReleaseAlibabaFootscanMiniReportFragmentFirstMtopResult(v *AlibabaFootscanMiniReportFragmentFirstMtopResult) {
	v.Message = ""
	v.Code = 0
	v.Data = nil
	poolAlibabaFootscanMiniReportFragmentFirstMtopResult.Put(v)
}
