package wdk

import (
	"sync"
)

// SeriesSortRequest 结构体
type SeriesSortRequest struct {
	// 有序行业属性对：行业属性key，属性值
	IndustryAttrList string `json:"industry_attr_list,omitempty" xml:"industry_attr_list,omitempty"`
	// 行业类型
	IndustryType string `json:"industry_type,omitempty" xml:"industry_type,omitempty"`
	// 系列编码
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
}

var poolSeriesSortRequest = sync.Pool{
	New: func() any {
		return new(SeriesSortRequest)
	},
}

// GetSeriesSortRequest() 从对象池中获取SeriesSortRequest
func GetSeriesSortRequest() *SeriesSortRequest {
	return poolSeriesSortRequest.Get().(*SeriesSortRequest)
}

// ReleaseSeriesSortRequest 释放SeriesSortRequest
func ReleaseSeriesSortRequest(v *SeriesSortRequest) {
	v.IndustryAttrList = ""
	v.IndustryType = ""
	v.SeriesId = 0
	poolSeriesSortRequest.Put(v)
}
