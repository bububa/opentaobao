package product

import (
	"sync"
)

// ItemSeriesRequest 结构体
type ItemSeriesRequest struct {
	// 市场
	Market string `json:"market,omitempty" xml:"market,omitempty"`
	// 系列名称
	SeriesName string `json:"series_name,omitempty" xml:"series_name,omitempty"`
	// 系列描述
	SeriesDesc string `json:"series_desc,omitempty" xml:"series_desc,omitempty"`
	// 品牌值,当品牌作为聚合属性时生效
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 系列id
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
	// 系列状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolItemSeriesRequest = sync.Pool{
	New: func() any {
		return new(ItemSeriesRequest)
	},
}

// GetItemSeriesRequest() 从对象池中获取ItemSeriesRequest
func GetItemSeriesRequest() *ItemSeriesRequest {
	return poolItemSeriesRequest.Get().(*ItemSeriesRequest)
}

// ReleaseItemSeriesRequest 释放ItemSeriesRequest
func ReleaseItemSeriesRequest(v *ItemSeriesRequest) {
	v.Market = ""
	v.SeriesName = ""
	v.SeriesDesc = ""
	v.BrandName = ""
	v.CatId = 0
	v.SeriesId = 0
	v.Status = 0
	poolItemSeriesRequest.Put(v)
}
