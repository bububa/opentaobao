package eticket

import (
	"sync"
)

// EticketTask 结构体
type EticketTask struct {
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolEticketTask = sync.Pool{
	New: func() any {
		return new(EticketTask)
	},
}

// GetEticketTask() 从对象池中获取EticketTask
func GetEticketTask() *EticketTask {
	return poolEticketTask.Get().(*EticketTask)
}

// ReleaseEticketTask 释放EticketTask
func ReleaseEticketTask(v *EticketTask) {
	v.OrderId = 0
	poolEticketTask.Put(v)
}
