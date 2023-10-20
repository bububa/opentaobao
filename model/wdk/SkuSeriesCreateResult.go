package wdk

import (
	"sync"
)

// SkuSeriesCreateResult 结构体
type SkuSeriesCreateResult struct {
	// 系列编码
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
}

var poolSkuSeriesCreateResult = sync.Pool{
	New: func() any {
		return new(SkuSeriesCreateResult)
	},
}

// GetSkuSeriesCreateResult() 从对象池中获取SkuSeriesCreateResult
func GetSkuSeriesCreateResult() *SkuSeriesCreateResult {
	return poolSkuSeriesCreateResult.Get().(*SkuSeriesCreateResult)
}

// ReleaseSkuSeriesCreateResult 释放SkuSeriesCreateResult
func ReleaseSkuSeriesCreateResult(v *SkuSeriesCreateResult) {
	v.SeriesId = 0
	poolSkuSeriesCreateResult.Put(v)
}
