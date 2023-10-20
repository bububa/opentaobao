package xhotel

import (
	"sync"
)

// TopAdsHtlDataQueryResult 结构体
type TopAdsHtlDataQueryResult struct {
	// topAdsHtlServiceDataList
	TopAdsHtlServiceDataList []Topadshtlservicedatalist `json:"top_ads_htl_service_data_list,omitempty" xml:"top_ads_htl_service_data_list>topadshtlservicedatalist,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolTopAdsHtlDataQueryResult = sync.Pool{
	New: func() any {
		return new(TopAdsHtlDataQueryResult)
	},
}

// GetTopAdsHtlDataQueryResult() 从对象池中获取TopAdsHtlDataQueryResult
func GetTopAdsHtlDataQueryResult() *TopAdsHtlDataQueryResult {
	return poolTopAdsHtlDataQueryResult.Get().(*TopAdsHtlDataQueryResult)
}

// ReleaseTopAdsHtlDataQueryResult 释放TopAdsHtlDataQueryResult
func ReleaseTopAdsHtlDataQueryResult(v *TopAdsHtlDataQueryResult) {
	v.TopAdsHtlServiceDataList = v.TopAdsHtlServiceDataList[:0]
	v.Count = 0
	poolTopAdsHtlDataQueryResult.Put(v)
}
