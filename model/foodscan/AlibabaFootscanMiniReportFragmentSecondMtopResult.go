package foodscan

import (
	"sync"
)

// AlibabaFootscanMiniReportFragmentSecondMtopResult 结构体
type AlibabaFootscanMiniReportFragmentSecondMtopResult struct {
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回i数据
	Data *AlibabaFootscanMiniReportFragmentSecondData `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaFootscanMiniReportFragmentSecondMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniReportFragmentSecondMtopResult)
	},
}

// GetAlibabaFootscanMiniReportFragmentSecondMtopResult() 从对象池中获取AlibabaFootscanMiniReportFragmentSecondMtopResult
func GetAlibabaFootscanMiniReportFragmentSecondMtopResult() *AlibabaFootscanMiniReportFragmentSecondMtopResult {
	return poolAlibabaFootscanMiniReportFragmentSecondMtopResult.Get().(*AlibabaFootscanMiniReportFragmentSecondMtopResult)
}

// ReleaseAlibabaFootscanMiniReportFragmentSecondMtopResult 释放AlibabaFootscanMiniReportFragmentSecondMtopResult
func ReleaseAlibabaFootscanMiniReportFragmentSecondMtopResult(v *AlibabaFootscanMiniReportFragmentSecondMtopResult) {
	v.Message = ""
	v.Code = 0
	v.Data = nil
	poolAlibabaFootscanMiniReportFragmentSecondMtopResult.Put(v)
}
