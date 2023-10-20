package alihealth2

import (
	"sync"
)

// IsvReserveRequest 结构体
type IsvReserveRequest struct {
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// isv 预约单ID
	SpReserveId string `json:"sp_reserve_id,omitempty" xml:"sp_reserve_id,omitempty"`
	// 预约时间
	ReserveTime string `json:"reserve_time,omitempty" xml:"reserve_time,omitempty"`
	// 预约单ID
	ReserveId int64 `json:"reserve_id,omitempty" xml:"reserve_id,omitempty"`
}

var poolIsvReserveRequest = sync.Pool{
	New: func() any {
		return new(IsvReserveRequest)
	},
}

// GetIsvReserveRequest() 从对象池中获取IsvReserveRequest
func GetIsvReserveRequest() *IsvReserveRequest {
	return poolIsvReserveRequest.Get().(*IsvReserveRequest)
}

// ReleaseIsvReserveRequest 释放IsvReserveRequest
func ReleaseIsvReserveRequest(v *IsvReserveRequest) {
	v.OrderId = ""
	v.SpReserveId = ""
	v.ReserveTime = ""
	v.ReserveId = 0
	poolIsvReserveRequest.Put(v)
}
