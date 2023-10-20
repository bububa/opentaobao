package tbtrade

import (
	"sync"
)

// OrderGroupResponse 结构体
type OrderGroupResponse struct {
	// 入参中的groupId
	GroupId string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 回传结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolOrderGroupResponse = sync.Pool{
	New: func() any {
		return new(OrderGroupResponse)
	},
}

// GetOrderGroupResponse() 从对象池中获取OrderGroupResponse
func GetOrderGroupResponse() *OrderGroupResponse {
	return poolOrderGroupResponse.Get().(*OrderGroupResponse)
}

// ReleaseOrderGroupResponse 释放OrderGroupResponse
func ReleaseOrderGroupResponse(v *OrderGroupResponse) {
	v.GroupId = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = false
	poolOrderGroupResponse.Put(v)
}
