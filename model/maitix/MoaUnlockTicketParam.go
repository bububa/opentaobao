package maitix

import (
	"sync"
)

// MoaUnlockTicketParam 结构体
type MoaUnlockTicketParam struct {
	// 大麦订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolMoaUnlockTicketParam = sync.Pool{
	New: func() any {
		return new(MoaUnlockTicketParam)
	},
}

// GetMoaUnlockTicketParam() 从对象池中获取MoaUnlockTicketParam
func GetMoaUnlockTicketParam() *MoaUnlockTicketParam {
	return poolMoaUnlockTicketParam.Get().(*MoaUnlockTicketParam)
}

// ReleaseMoaUnlockTicketParam 释放MoaUnlockTicketParam
func ReleaseMoaUnlockTicketParam(v *MoaUnlockTicketParam) {
	v.OrderId = 0
	poolMoaUnlockTicketParam.Put(v)
}
