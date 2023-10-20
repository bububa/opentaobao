package omniorder

import (
	"sync"
)

// Door2doorConsumeRequest 结构体
type Door2doorConsumeRequest struct {
	// 核销码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 淘宝主订单ID
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}

var poolDoor2doorConsumeRequest = sync.Pool{
	New: func() any {
		return new(Door2doorConsumeRequest)
	},
}

// GetDoor2doorConsumeRequest() 从对象池中获取Door2doorConsumeRequest
func GetDoor2doorConsumeRequest() *Door2doorConsumeRequest {
	return poolDoor2doorConsumeRequest.Get().(*Door2doorConsumeRequest)
}

// ReleaseDoor2doorConsumeRequest 释放Door2doorConsumeRequest
func ReleaseDoor2doorConsumeRequest(v *Door2doorConsumeRequest) {
	v.Code = ""
	v.Operator = ""
	v.MainOrderId = 0
	poolDoor2doorConsumeRequest.Put(v)
}
