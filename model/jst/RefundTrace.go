package jst

import (
	"sync"
)

// RefundTrace 结构体
type RefundTrace struct {
	// 动作发生的时间
	ActionTime string `json:"action_time,omitempty" xml:"action_time,omitempty"`
	// 应用标题
	AppTitle string `json:"app_title,omitempty" xml:"app_title,omitempty"`
	// 备注字段
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 回流的订单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 退款编号
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 交易tid
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolRefundTrace = sync.Pool{
	New: func() any {
		return new(RefundTrace)
	},
}

// GetRefundTrace() 从对象池中获取RefundTrace
func GetRefundTrace() *RefundTrace {
	return poolRefundTrace.Get().(*RefundTrace)
}

// ReleaseRefundTrace 释放RefundTrace
func ReleaseRefundTrace(v *RefundTrace) {
	v.ActionTime = ""
	v.AppTitle = ""
	v.Remark = ""
	v.Status = ""
	v.RefundId = 0
	v.Tid = 0
	poolRefundTrace.Put(v)
}
