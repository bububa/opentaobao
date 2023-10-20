package wdk

import (
	"sync"
)

// SkuSeriesEditRequest 结构体
type SkuSeriesEditRequest struct {
	// 系列描述
	SeriesDesc string `json:"series_desc,omitempty" xml:"series_desc,omitempty"`
	// 系列名称
	SeriesName string `json:"series_name,omitempty" xml:"series_name,omitempty"`
	// 系列编码
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
}

var poolSkuSeriesEditRequest = sync.Pool{
	New: func() any {
		return new(SkuSeriesEditRequest)
	},
}

// GetSkuSeriesEditRequest() 从对象池中获取SkuSeriesEditRequest
func GetSkuSeriesEditRequest() *SkuSeriesEditRequest {
	return poolSkuSeriesEditRequest.Get().(*SkuSeriesEditRequest)
}

// ReleaseSkuSeriesEditRequest 释放SkuSeriesEditRequest
func ReleaseSkuSeriesEditRequest(v *SkuSeriesEditRequest) {
	v.SeriesDesc = ""
	v.SeriesName = ""
	v.SeriesId = 0
	poolSkuSeriesEditRequest.Put(v)
}
