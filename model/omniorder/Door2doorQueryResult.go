package omniorder

import (
	"sync"
)

// Door2doorQueryResult 结构体
type Door2doorQueryResult struct {
	// 码对应的淘宝主订单ID
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}

var poolDoor2doorQueryResult = sync.Pool{
	New: func() any {
		return new(Door2doorQueryResult)
	},
}

// GetDoor2doorQueryResult() 从对象池中获取Door2doorQueryResult
func GetDoor2doorQueryResult() *Door2doorQueryResult {
	return poolDoor2doorQueryResult.Get().(*Door2doorQueryResult)
}

// ReleaseDoor2doorQueryResult 释放Door2doorQueryResult
func ReleaseDoor2doorQueryResult(v *Door2doorQueryResult) {
	v.MainOrderId = 0
	poolDoor2doorQueryResult.Put(v)
}
