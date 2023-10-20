package wdk

import (
	"sync"
)

// TimeSliceGetRequest 结构体
type TimeSliceGetRequest struct {
	// 购买商品列表
	SkuList []PromiseSkuInfo `json:"sku_list,omitempty" xml:"sku_list>promise_sku_info,omitempty"`
	// 渠道店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 配送地址经纬度
	Geo string `json:"geo,omitempty" xml:"geo,omitempty"`
}

var poolTimeSliceGetRequest = sync.Pool{
	New: func() any {
		return new(TimeSliceGetRequest)
	},
}

// GetTimeSliceGetRequest() 从对象池中获取TimeSliceGetRequest
func GetTimeSliceGetRequest() *TimeSliceGetRequest {
	return poolTimeSliceGetRequest.Get().(*TimeSliceGetRequest)
}

// ReleaseTimeSliceGetRequest 释放TimeSliceGetRequest
func ReleaseTimeSliceGetRequest(v *TimeSliceGetRequest) {
	v.SkuList = v.SkuList[:0]
	v.ShopId = ""
	v.Geo = ""
	poolTimeSliceGetRequest.Put(v)
}
