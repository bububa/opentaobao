package maitix

import (
	"sync"
)

// UnLockTicketResponse 结构体
type UnLockTicketResponse struct {
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolUnLockTicketResponse = sync.Pool{
	New: func() any {
		return new(UnLockTicketResponse)
	},
}

// GetUnLockTicketResponse() 从对象池中获取UnLockTicketResponse
func GetUnLockTicketResponse() *UnLockTicketResponse {
	return poolUnLockTicketResponse.Get().(*UnLockTicketResponse)
}

// ReleaseUnLockTicketResponse 释放UnLockTicketResponse
func ReleaseUnLockTicketResponse(v *UnLockTicketResponse) {
	v.OrderId = 0
	poolUnLockTicketResponse.Put(v)
}
