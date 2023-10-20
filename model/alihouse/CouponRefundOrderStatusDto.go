package alihouse

import (
	"sync"
)

// CouponRefundOrderStatusDto 结构体
type CouponRefundOrderStatusDto struct {
	// 截止时间
	CancelDeadlineTime string `json:"cancel_deadline_time,omitempty" xml:"cancel_deadline_time,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 取消原因
	CancelRemark string `json:"cancel_remark,omitempty" xml:"cancel_remark,omitempty"`
	// 取消完成时间
	CancelFinishedTime string `json:"cancel_finished_time,omitempty" xml:"cancel_finished_time,omitempty"`
	// 节点
	CancelNode int64 `json:"cancel_node,omitempty" xml:"cancel_node,omitempty"`
	// 角色
	OperatorType int64 `json:"operator_type,omitempty" xml:"operator_type,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 淘宝退款单id
	TbRefundOrderId int64 `json:"tb_refund_order_id,omitempty" xml:"tb_refund_order_id,omitempty"`
	// 退款单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolCouponRefundOrderStatusDto = sync.Pool{
	New: func() any {
		return new(CouponRefundOrderStatusDto)
	},
}

// GetCouponRefundOrderStatusDto() 从对象池中获取CouponRefundOrderStatusDto
func GetCouponRefundOrderStatusDto() *CouponRefundOrderStatusDto {
	return poolCouponRefundOrderStatusDto.Get().(*CouponRefundOrderStatusDto)
}

// ReleaseCouponRefundOrderStatusDto 释放CouponRefundOrderStatusDto
func ReleaseCouponRefundOrderStatusDto(v *CouponRefundOrderStatusDto) {
	v.CancelDeadlineTime = ""
	v.StartTime = ""
	v.StatusDesc = ""
	v.CancelRemark = ""
	v.CancelFinishedTime = ""
	v.CancelNode = 0
	v.OperatorType = 0
	v.Status = 0
	v.TbRefundOrderId = 0
	v.Id = 0
	poolCouponRefundOrderStatusDto.Put(v)
}
