package product

import (
	"sync"
)

// SeriesItemRequest 结构体
type SeriesItemRequest struct {
	// 市场
	Market string `json:"market,omitempty" xml:"market,omitempty"`
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 分组名称
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// 定制展示名称
	CustomVersion string `json:"custom_version,omitempty" xml:"custom_version,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 系列id
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
}

var poolSeriesItemRequest = sync.Pool{
	New: func() any {
		return new(SeriesItemRequest)
	},
}

// GetSeriesItemRequest() 从对象池中获取SeriesItemRequest
func GetSeriesItemRequest() *SeriesItemRequest {
	return poolSeriesItemRequest.Get().(*SeriesItemRequest)
}

// ReleaseSeriesItemRequest 释放SeriesItemRequest
func ReleaseSeriesItemRequest(v *SeriesItemRequest) {
	v.Market = ""
	v.ItemId = ""
	v.GroupName = ""
	v.CustomVersion = ""
	v.Sort = 0
	v.SeriesId = 0
	poolSeriesItemRequest.Put(v)
}
