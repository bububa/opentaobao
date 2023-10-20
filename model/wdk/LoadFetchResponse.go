package wdk

import (
	"sync"
)

// LoadFetchResponse 结构体
type LoadFetchResponse struct {
	// 取货单list
	FetchAggregates []FetchAggregateSdo `json:"fetch_aggregates,omitempty" xml:"fetch_aggregates>fetch_aggregate_sdo,omitempty"`
	// 退货取货单ID
	FetchOrderId string `json:"fetch_order_id,omitempty" xml:"fetch_order_id,omitempty"`
}

var poolLoadFetchResponse = sync.Pool{
	New: func() any {
		return new(LoadFetchResponse)
	},
}

// GetLoadFetchResponse() 从对象池中获取LoadFetchResponse
func GetLoadFetchResponse() *LoadFetchResponse {
	return poolLoadFetchResponse.Get().(*LoadFetchResponse)
}

// ReleaseLoadFetchResponse 释放LoadFetchResponse
func ReleaseLoadFetchResponse(v *LoadFetchResponse) {
	v.FetchAggregates = v.FetchAggregates[:0]
	v.FetchOrderId = ""
	poolLoadFetchResponse.Put(v)
}
