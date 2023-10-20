package maitix

import (
	"sync"
)

// LockTicketResponse 结构体
type LockTicketResponse struct {
	// 子订单列表
	SubOrderDtos []LockTicketSubOrderDto `json:"sub_order_dtos,omitempty" xml:"sub_order_dtos>lock_ticket_sub_order_dto,omitempty"`
	// 大麦订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 总价
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 快递费
	ExpressFee int64 `json:"express_fee,omitempty" xml:"express_fee,omitempty"`
}

var poolLockTicketResponse = sync.Pool{
	New: func() any {
		return new(LockTicketResponse)
	},
}

// GetLockTicketResponse() 从对象池中获取LockTicketResponse
func GetLockTicketResponse() *LockTicketResponse {
	return poolLockTicketResponse.Get().(*LockTicketResponse)
}

// ReleaseLockTicketResponse 释放LockTicketResponse
func ReleaseLockTicketResponse(v *LockTicketResponse) {
	v.SubOrderDtos = v.SubOrderDtos[:0]
	v.OrderId = ""
	v.TotalAmount = 0
	v.ExpressFee = 0
	poolLockTicketResponse.Put(v)
}
