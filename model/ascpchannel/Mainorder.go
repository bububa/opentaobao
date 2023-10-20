package ascpchannel

import (
	"sync"
)

// Mainorder 结构体
type Mainorder struct {
	// 操作id
	OperationOrderId string `json:"operation_order_id,omitempty" xml:"operation_order_id,omitempty"`
	// 商家Uic_id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolMainorder = sync.Pool{
	New: func() any {
		return new(Mainorder)
	},
}

// GetMainorder() 从对象池中获取Mainorder
func GetMainorder() *Mainorder {
	return poolMainorder.Get().(*Mainorder)
}

// ReleaseMainorder 释放Mainorder
func ReleaseMainorder(v *Mainorder) {
	v.OperationOrderId = ""
	v.UserId = 0
	poolMainorder.Put(v)
}
