package ieagency

import (
	"sync"
)

// RefundOrderDetailVo 结构体
type RefundOrderDetailVo struct {
	// 商家回复
	AgreeAnswer string `json:"agree_answer,omitempty" xml:"agree_answer,omitempty"`
	// 同意时间
	AgreeTime string `json:"agree_time,omitempty" xml:"agree_time,omitempty"`
	// 最晚处理时间
	LatestProcessTime string `json:"latest_process_time,omitempty" xml:"latest_process_time,omitempty"`
	// 接受回复
	ReceiveAnswer string `json:"receive_answer,omitempty" xml:"receive_answer,omitempty"`
	// 接受退票时间
	ReceiveTime string `json:"receive_time,omitempty" xml:"receive_time,omitempty"`
	// 拒绝答复
	RefuseAnswer string `json:"refuse_answer,omitempty" xml:"refuse_answer,omitempty"`
	// 拒绝时间
	RefuseTime string `json:"refuse_time,omitempty" xml:"refuse_time,omitempty"`
	// 创建时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 模型版本
	ModelVersion string `json:"model_version,omitempty" xml:"model_version,omitempty"`
	// 极速类型
	InstantType int64 `json:"instant_type,omitempty" xml:"instant_type,omitempty"`
	// 多次退次数
	MultiRefundIndex int64 `json:"multi_refund_index,omitempty" xml:"multi_refund_index,omitempty"`
	// 申请单退商品粒度
	RefundTicketDimension int64 `json:"refund_ticket_dimension,omitempty" xml:"refund_ticket_dimension,omitempty"`
	// 是否补退
	BuiTui bool `json:"bui_tui,omitempty" xml:"bui_tui,omitempty"`
	// 限时免费退
	TimeLimitRefund bool `json:"time_limit_refund,omitempty" xml:"time_limit_refund,omitempty"`
}

var poolRefundOrderDetailVo = sync.Pool{
	New: func() any {
		return new(RefundOrderDetailVo)
	},
}

// GetRefundOrderDetailVo() 从对象池中获取RefundOrderDetailVo
func GetRefundOrderDetailVo() *RefundOrderDetailVo {
	return poolRefundOrderDetailVo.Get().(*RefundOrderDetailVo)
}

// ReleaseRefundOrderDetailVo 释放RefundOrderDetailVo
func ReleaseRefundOrderDetailVo(v *RefundOrderDetailVo) {
	v.AgreeAnswer = ""
	v.AgreeTime = ""
	v.LatestProcessTime = ""
	v.ReceiveAnswer = ""
	v.ReceiveTime = ""
	v.RefuseAnswer = ""
	v.RefuseTime = ""
	v.ApplyTime = ""
	v.ModelVersion = ""
	v.InstantType = 0
	v.MultiRefundIndex = 0
	v.RefundTicketDimension = 0
	v.BuiTui = false
	v.TimeLimitRefund = false
	poolRefundOrderDetailVo.Put(v)
}
