package wdk

import (
	"sync"
)

// ItemDiscountActivityQueryRequest 结构体
type ItemDiscountActivityQueryRequest struct {
	// erp外部活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}

var poolItemDiscountActivityQueryRequest = sync.Pool{
	New: func() any {
		return new(ItemDiscountActivityQueryRequest)
	},
}

// GetItemDiscountActivityQueryRequest() 从对象池中获取ItemDiscountActivityQueryRequest
func GetItemDiscountActivityQueryRequest() *ItemDiscountActivityQueryRequest {
	return poolItemDiscountActivityQueryRequest.Get().(*ItemDiscountActivityQueryRequest)
}

// ReleaseItemDiscountActivityQueryRequest 释放ItemDiscountActivityQueryRequest
func ReleaseItemDiscountActivityQueryRequest(v *ItemDiscountActivityQueryRequest) {
	v.OutActId = ""
	v.ActId = 0
	poolItemDiscountActivityQueryRequest.Put(v)
}
