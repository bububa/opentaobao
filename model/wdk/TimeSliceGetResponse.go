package wdk

import (
	"sync"
)

// TimeSliceGetResponse 结构体
type TimeSliceGetResponse struct {
	// 时间片列表
	TimeSliceList []PromiseTimeSlice `json:"time_slice_list,omitempty" xml:"time_slice_list>promise_time_slice,omitempty"`
	// 商品信息
	ProductList []FulfillProduct `json:"product_list,omitempty" xml:"product_list>fulfill_product,omitempty"`
}

var poolTimeSliceGetResponse = sync.Pool{
	New: func() any {
		return new(TimeSliceGetResponse)
	},
}

// GetTimeSliceGetResponse() 从对象池中获取TimeSliceGetResponse
func GetTimeSliceGetResponse() *TimeSliceGetResponse {
	return poolTimeSliceGetResponse.Get().(*TimeSliceGetResponse)
}

// ReleaseTimeSliceGetResponse 释放TimeSliceGetResponse
func ReleaseTimeSliceGetResponse(v *TimeSliceGetResponse) {
	v.TimeSliceList = v.TimeSliceList[:0]
	v.ProductList = v.ProductList[:0]
	poolTimeSliceGetResponse.Put(v)
}
