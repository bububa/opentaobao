package wdk

import (
	"sync"
)

// LoadFetchReq 结构体
type LoadFetchReq struct {
	// 取货单id
	FetchOrderId string `json:"fetch_order_id,omitempty" xml:"fetch_order_id,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 主单号
	MainOutOrderId string `json:"main_out_order_id,omitempty" xml:"main_out_order_id,omitempty"`
}

var poolLoadFetchReq = sync.Pool{
	New: func() any {
		return new(LoadFetchReq)
	},
}

// GetLoadFetchReq() 从对象池中获取LoadFetchReq
func GetLoadFetchReq() *LoadFetchReq {
	return poolLoadFetchReq.Get().(*LoadFetchReq)
}

// ReleaseLoadFetchReq 释放LoadFetchReq
func ReleaseLoadFetchReq(v *LoadFetchReq) {
	v.FetchOrderId = ""
	v.StoreId = ""
	v.MainOutOrderId = ""
	poolLoadFetchReq.Put(v)
}
