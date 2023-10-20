package tbrefund

import (
	"sync"
)

// RefundRemindTimeout 结构体
type RefundRemindTimeout struct {
	// 超时时间。格式:yyyy-MM-dd HH:mm:ss
	Timeout string `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 提醒的类型（退款详情中提示信息的类型）
	RemindType int64 `json:"remind_type,omitempty" xml:"remind_type,omitempty"`
	// 是否存在超时。可选值:true(是),false(否)
	ExistTimeout bool `json:"exist_timeout,omitempty" xml:"exist_timeout,omitempty"`
}

var poolRefundRemindTimeout = sync.Pool{
	New: func() any {
		return new(RefundRemindTimeout)
	},
}

// GetRefundRemindTimeout() 从对象池中获取RefundRemindTimeout
func GetRefundRemindTimeout() *RefundRemindTimeout {
	return poolRefundRemindTimeout.Get().(*RefundRemindTimeout)
}

// ReleaseRefundRemindTimeout 释放RefundRemindTimeout
func ReleaseRefundRemindTimeout(v *RefundRemindTimeout) {
	v.Timeout = ""
	v.RemindType = 0
	v.ExistTimeout = false
	poolRefundRemindTimeout.Put(v)
}
