package tmallcar

import (
	"sync"
)

// SyncInfoReq 结构体
type SyncInfoReq struct {
	// 幂等id
	Idempotent string `json:"idempotent,omitempty" xml:"idempotent,omitempty"`
	// 变更时间
	ChangeTime string `json:"change_time,omitempty" xml:"change_time,omitempty"`
	// 交付状态详细描述
	DeliveryStatusDesc string `json:"delivery_status_desc,omitempty" xml:"delivery_status_desc,omitempty"`
	// 交付状态
	DeliveryStatus string `json:"delivery_status,omitempty" xml:"delivery_status,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 物流轨迹
	LogisticsTraceReq *LogisticsTraceReq `json:"logistics_trace_req,omitempty" xml:"logistics_trace_req,omitempty"`
}

var poolSyncInfoReq = sync.Pool{
	New: func() any {
		return new(SyncInfoReq)
	},
}

// GetSyncInfoReq() 从对象池中获取SyncInfoReq
func GetSyncInfoReq() *SyncInfoReq {
	return poolSyncInfoReq.Get().(*SyncInfoReq)
}

// ReleaseSyncInfoReq 释放SyncInfoReq
func ReleaseSyncInfoReq(v *SyncInfoReq) {
	v.Idempotent = ""
	v.ChangeTime = ""
	v.DeliveryStatusDesc = ""
	v.DeliveryStatus = ""
	v.OrderId = 0
	v.LogisticsTraceReq = nil
	poolSyncInfoReq.Put(v)
}
