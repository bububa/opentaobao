package train

import (
	"sync"
)

// LockOrderRq 结构体
type LockOrderRq struct {
	// 票信息
	Tickets []TicketDto `json:"tickets,omitempty" xml:"tickets>ticket_dto,omitempty"`
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}

var poolLockOrderRq = sync.Pool{
	New: func() any {
		return new(LockOrderRq)
	},
}

// GetLockOrderRq() 从对象池中获取LockOrderRq
func GetLockOrderRq() *LockOrderRq {
	return poolLockOrderRq.Get().(*LockOrderRq)
}

// ReleaseLockOrderRq 释放LockOrderRq
func ReleaseLockOrderRq(v *LockOrderRq) {
	v.Tickets = v.Tickets[:0]
	v.TtpOrderId = 0
	poolLockOrderRq.Put(v)
}
