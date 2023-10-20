package tbtrade

import (
	"sync"
)

// AgreeRefundCheck 结构体
type AgreeRefundCheck struct {
	// 提示文案
	DeliveryTips string `json:"delivery_tips,omitempty" xml:"delivery_tips,omitempty"`
	// 流程状态
	DeliveryProcess string `json:"delivery_process,omitempty" xml:"delivery_process,omitempty"`
	// 订单id
	DetailOrderId string `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
}

var poolAgreeRefundCheck = sync.Pool{
	New: func() any {
		return new(AgreeRefundCheck)
	},
}

// GetAgreeRefundCheck() 从对象池中获取AgreeRefundCheck
func GetAgreeRefundCheck() *AgreeRefundCheck {
	return poolAgreeRefundCheck.Get().(*AgreeRefundCheck)
}

// ReleaseAgreeRefundCheck 释放AgreeRefundCheck
func ReleaseAgreeRefundCheck(v *AgreeRefundCheck) {
	v.DeliveryTips = ""
	v.DeliveryProcess = ""
	v.DetailOrderId = ""
	poolAgreeRefundCheck.Put(v)
}
