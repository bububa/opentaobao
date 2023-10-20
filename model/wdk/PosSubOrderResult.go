package wdk

import (
	"sync"
)

// PosSubOrderResult 结构体
type PosSubOrderResult struct {
	// outOrderId
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// subOrderId
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

var poolPosSubOrderResult = sync.Pool{
	New: func() any {
		return new(PosSubOrderResult)
	},
}

// GetPosSubOrderResult() 从对象池中获取PosSubOrderResult
func GetPosSubOrderResult() *PosSubOrderResult {
	return poolPosSubOrderResult.Get().(*PosSubOrderResult)
}

// ReleasePosSubOrderResult 释放PosSubOrderResult
func ReleasePosSubOrderResult(v *PosSubOrderResult) {
	v.OutOrderId = ""
	v.SubOrderId = 0
	poolPosSubOrderResult.Put(v)
}
