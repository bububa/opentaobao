package foodscan

import (
	"sync"
)

// AlibabaFootscanMiniQueryMobilereportMtopResult 结构体
type AlibabaFootscanMiniQueryMobilereportMtopResult struct {
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回i数据
	Data *AlibabaFootscanMiniQueryMobilereportData `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaFootscanMiniQueryMobilereportMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniQueryMobilereportMtopResult)
	},
}

// GetAlibabaFootscanMiniQueryMobilereportMtopResult() 从对象池中获取AlibabaFootscanMiniQueryMobilereportMtopResult
func GetAlibabaFootscanMiniQueryMobilereportMtopResult() *AlibabaFootscanMiniQueryMobilereportMtopResult {
	return poolAlibabaFootscanMiniQueryMobilereportMtopResult.Get().(*AlibabaFootscanMiniQueryMobilereportMtopResult)
}

// ReleaseAlibabaFootscanMiniQueryMobilereportMtopResult 释放AlibabaFootscanMiniQueryMobilereportMtopResult
func ReleaseAlibabaFootscanMiniQueryMobilereportMtopResult(v *AlibabaFootscanMiniQueryMobilereportMtopResult) {
	v.Message = ""
	v.Code = 0
	v.Data = nil
	poolAlibabaFootscanMiniQueryMobilereportMtopResult.Put(v)
}
