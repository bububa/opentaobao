package wdk

import (
	"sync"
)

// CreateReverseResponse 结构体
type CreateReverseResponse struct {
	// 外部单号
	OutBizOrderIds []string `json:"out_biz_order_ids,omitempty" xml:"out_biz_order_ids>string,omitempty"`
	// 逆向单ids
	ReverseIds []int64 `json:"reverse_ids,omitempty" xml:"reverse_ids>int64,omitempty"`
	// tp单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolCreateReverseResponse = sync.Pool{
	New: func() any {
		return new(CreateReverseResponse)
	},
}

// GetCreateReverseResponse() 从对象池中获取CreateReverseResponse
func GetCreateReverseResponse() *CreateReverseResponse {
	return poolCreateReverseResponse.Get().(*CreateReverseResponse)
}

// ReleaseCreateReverseResponse 释放CreateReverseResponse
func ReleaseCreateReverseResponse(v *CreateReverseResponse) {
	v.OutBizOrderIds = v.OutBizOrderIds[:0]
	v.ReverseIds = v.ReverseIds[:0]
	v.OutOrderId = ""
	v.RequestId = ""
	v.StoreId = ""
	poolCreateReverseResponse.Put(v)
}
