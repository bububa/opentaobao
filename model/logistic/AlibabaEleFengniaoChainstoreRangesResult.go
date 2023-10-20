package logistic

import (
	"sync"
)

// AlibabaEleFengniaoChainstoreRangesResult 结构体
type AlibabaEleFengniaoChainstoreRangesResult struct {
	// ranges
	Ranges []Range `json:"ranges,omitempty" xml:"ranges>range,omitempty"`
	// 到达圈标识
	RangeId int64 `json:"range_id,omitempty" xml:"range_id,omitempty"`
	// 配送圈类型 1-日间 2- 晚上  3-全天
	RangeType int64 `json:"range_type,omitempty" xml:"range_type,omitempty"`
}

var poolAlibabaEleFengniaoChainstoreRangesResult = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoChainstoreRangesResult)
	},
}

// GetAlibabaEleFengniaoChainstoreRangesResult() 从对象池中获取AlibabaEleFengniaoChainstoreRangesResult
func GetAlibabaEleFengniaoChainstoreRangesResult() *AlibabaEleFengniaoChainstoreRangesResult {
	return poolAlibabaEleFengniaoChainstoreRangesResult.Get().(*AlibabaEleFengniaoChainstoreRangesResult)
}

// ReleaseAlibabaEleFengniaoChainstoreRangesResult 释放AlibabaEleFengniaoChainstoreRangesResult
func ReleaseAlibabaEleFengniaoChainstoreRangesResult(v *AlibabaEleFengniaoChainstoreRangesResult) {
	v.Ranges = v.Ranges[:0]
	v.RangeId = 0
	v.RangeType = 0
	poolAlibabaEleFengniaoChainstoreRangesResult.Put(v)
}
