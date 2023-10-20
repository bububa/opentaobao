package westcrm

import (
	"sync"
)

// UserConsumerVo 结构体
type UserConsumerVo struct {
	// 订单
	Orders []OrderVo `json:"orders,omitempty" xml:"orders>order_vo,omitempty"`
	// 消费总额
	TotalPay string `json:"total_pay,omitempty" xml:"total_pay,omitempty"`
	// 用户id
	IbUserId int64 `json:"ib_user_id,omitempty" xml:"ib_user_id,omitempty"`
}

var poolUserConsumerVo = sync.Pool{
	New: func() any {
		return new(UserConsumerVo)
	},
}

// GetUserConsumerVo() 从对象池中获取UserConsumerVo
func GetUserConsumerVo() *UserConsumerVo {
	return poolUserConsumerVo.Get().(*UserConsumerVo)
}

// ReleaseUserConsumerVo 释放UserConsumerVo
func ReleaseUserConsumerVo(v *UserConsumerVo) {
	v.Orders = v.Orders[:0]
	v.TotalPay = ""
	v.IbUserId = 0
	poolUserConsumerVo.Put(v)
}
