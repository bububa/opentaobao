package wdk

import (
	"sync"
)

// QueryTimeSliceReq 结构体
type QueryTimeSliceReq struct {
	// 子单list
	SubOutOrderIds []string `json:"sub_out_order_ids,omitempty" xml:"sub_out_order_ids>string,omitempty"`
	// 主单号
	MainOutOrderId string `json:"main_out_order_id,omitempty" xml:"main_out_order_id,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolQueryTimeSliceReq = sync.Pool{
	New: func() any {
		return new(QueryTimeSliceReq)
	},
}

// GetQueryTimeSliceReq() 从对象池中获取QueryTimeSliceReq
func GetQueryTimeSliceReq() *QueryTimeSliceReq {
	return poolQueryTimeSliceReq.Get().(*QueryTimeSliceReq)
}

// ReleaseQueryTimeSliceReq 释放QueryTimeSliceReq
func ReleaseQueryTimeSliceReq(v *QueryTimeSliceReq) {
	v.SubOutOrderIds = v.SubOutOrderIds[:0]
	v.MainOutOrderId = ""
	v.StoreId = ""
	poolQueryTimeSliceReq.Put(v)
}
