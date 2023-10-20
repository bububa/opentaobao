package wdk

import (
	"sync"
)

// SkuSeriesEditResult 结构体
type SkuSeriesEditResult struct {
	// 成功的商品编码集
	FailedSkuCodes []string `json:"failed_sku_codes,omitempty" xml:"failed_sku_codes>string,omitempty"`
	// 失败的商品编码集
	SuccessedSkuCodes []string `json:"successed_sku_codes,omitempty" xml:"successed_sku_codes>string,omitempty"`
	// 系列编码
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
}

var poolSkuSeriesEditResult = sync.Pool{
	New: func() any {
		return new(SkuSeriesEditResult)
	},
}

// GetSkuSeriesEditResult() 从对象池中获取SkuSeriesEditResult
func GetSkuSeriesEditResult() *SkuSeriesEditResult {
	return poolSkuSeriesEditResult.Get().(*SkuSeriesEditResult)
}

// ReleaseSkuSeriesEditResult 释放SkuSeriesEditResult
func ReleaseSkuSeriesEditResult(v *SkuSeriesEditResult) {
	v.FailedSkuCodes = v.FailedSkuCodes[:0]
	v.SuccessedSkuCodes = v.SuccessedSkuCodes[:0]
	v.SeriesId = 0
	poolSkuSeriesEditResult.Put(v)
}
