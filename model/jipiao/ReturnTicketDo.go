package jipiao

import (
	"sync"
)

// ReturnTicketDo 结构体
type ReturnTicketDo struct {
	// 申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 申请单ID
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 申请单状态（1初始 2接受 3确认 4失败 5买家处理 6卖家处理 7等待小二回填 8退款成功）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolReturnTicketDo = sync.Pool{
	New: func() any {
		return new(ReturnTicketDo)
	},
}

// GetReturnTicketDo() 从对象池中获取ReturnTicketDo
func GetReturnTicketDo() *ReturnTicketDo {
	return poolReturnTicketDo.Get().(*ReturnTicketDo)
}

// ReleaseReturnTicketDo 释放ReturnTicketDo
func ReleaseReturnTicketDo(v *ReturnTicketDo) {
	v.ApplyTime = ""
	v.ApplyId = 0
	v.OrderId = 0
	v.Status = 0
	poolReturnTicketDo.Put(v)
}
