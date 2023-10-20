package maitix

import (
	"sync"
)

// MoaConfirmOrderParam 结构体
type MoaConfirmOrderParam struct {
	// 大麦订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolMoaConfirmOrderParam = sync.Pool{
	New: func() any {
		return new(MoaConfirmOrderParam)
	},
}

// GetMoaConfirmOrderParam() 从对象池中获取MoaConfirmOrderParam
func GetMoaConfirmOrderParam() *MoaConfirmOrderParam {
	return poolMoaConfirmOrderParam.Get().(*MoaConfirmOrderParam)
}

// ReleaseMoaConfirmOrderParam 释放MoaConfirmOrderParam
func ReleaseMoaConfirmOrderParam(v *MoaConfirmOrderParam) {
	v.OrderId = 0
	poolMoaConfirmOrderParam.Put(v)
}
