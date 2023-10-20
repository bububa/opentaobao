package xhotel

import (
	"sync"
)

// TopAdsTripSvcQueryResult 结构体
type TopAdsTripSvcQueryResult struct {
	// topAdsHtlServiceDataList
	TopAdsHtlServiceDataList []Topadshtlservicedatalist `json:"top_ads_htl_service_data_list,omitempty" xml:"top_ads_htl_service_data_list>topadshtlservicedatalist,omitempty"`
	// count
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolTopAdsTripSvcQueryResult = sync.Pool{
	New: func() any {
		return new(TopAdsTripSvcQueryResult)
	},
}

// GetTopAdsTripSvcQueryResult() 从对象池中获取TopAdsTripSvcQueryResult
func GetTopAdsTripSvcQueryResult() *TopAdsTripSvcQueryResult {
	return poolTopAdsTripSvcQueryResult.Get().(*TopAdsTripSvcQueryResult)
}

// ReleaseTopAdsTripSvcQueryResult 释放TopAdsTripSvcQueryResult
func ReleaseTopAdsTripSvcQueryResult(v *TopAdsTripSvcQueryResult) {
	v.TopAdsHtlServiceDataList = v.TopAdsHtlServiceDataList[:0]
	v.Count = 0
	poolTopAdsTripSvcQueryResult.Put(v)
}
