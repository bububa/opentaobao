package car

import (
	"sync"
)

// CrsOrderCompleteParam 结构体
type CrsOrderCompleteParam struct {
	// 飞猪订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolCrsOrderCompleteParam = sync.Pool{
	New: func() any {
		return new(CrsOrderCompleteParam)
	},
}

// GetCrsOrderCompleteParam() 从对象池中获取CrsOrderCompleteParam
func GetCrsOrderCompleteParam() *CrsOrderCompleteParam {
	return poolCrsOrderCompleteParam.Get().(*CrsOrderCompleteParam)
}

// ReleaseCrsOrderCompleteParam 释放CrsOrderCompleteParam
func ReleaseCrsOrderCompleteParam(v *CrsOrderCompleteParam) {
	v.OrderId = 0
	poolCrsOrderCompleteParam.Put(v)
}
