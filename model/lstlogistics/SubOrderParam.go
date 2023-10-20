package lstlogistics

import (
	"sync"
)

// SubOrderParam 结构体
type SubOrderParam struct {
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 购买数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolSubOrderParam = sync.Pool{
	New: func() any {
		return new(SubOrderParam)
	},
}

// GetSubOrderParam() 从对象池中获取SubOrderParam
func GetSubOrderParam() *SubOrderParam {
	return poolSubOrderParam.Get().(*SubOrderParam)
}

// ReleaseSubOrderParam 释放SubOrderParam
func ReleaseSubOrderParam(v *SubOrderParam) {
	v.SubOrderId = 0
	v.Amount = 0
	poolSubOrderParam.Put(v)
}
