package flightuppc

import (
	"sync"
)

// InsReverseOrderReq 结构体
type InsReverseOrderReq struct {
	// 保险信息列表
	Insureds []InsPersonAndAirSegmentDto `json:"insureds,omitempty" xml:"insureds>ins_person_and_air_segment_dto,omitempty"`
	// 保险订单号
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
}

var poolInsReverseOrderReq = sync.Pool{
	New: func() any {
		return new(InsReverseOrderReq)
	},
}

// GetInsReverseOrderReq() 从对象池中获取InsReverseOrderReq
func GetInsReverseOrderReq() *InsReverseOrderReq {
	return poolInsReverseOrderReq.Get().(*InsReverseOrderReq)
}

// ReleaseInsReverseOrderReq 释放InsReverseOrderReq
func ReleaseInsReverseOrderReq(v *InsReverseOrderReq) {
	v.Insureds = v.Insureds[:0]
	v.TpOrderId = 0
	poolInsReverseOrderReq.Put(v)
}
