package ieagency

import (
	"sync"
)

// IeBookPayOrderVo 结构体
type IeBookPayOrderVo struct {
	// orderId
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolIeBookPayOrderVo = sync.Pool{
	New: func() any {
		return new(IeBookPayOrderVo)
	},
}

// GetIeBookPayOrderVo() 从对象池中获取IeBookPayOrderVo
func GetIeBookPayOrderVo() *IeBookPayOrderVo {
	return poolIeBookPayOrderVo.Get().(*IeBookPayOrderVo)
}

// ReleaseIeBookPayOrderVo 释放IeBookPayOrderVo
func ReleaseIeBookPayOrderVo(v *IeBookPayOrderVo) {
	v.OrderId = 0
	poolIeBookPayOrderVo.Put(v)
}
