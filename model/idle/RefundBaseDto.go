package idle

import (
	"sync"
)

// RefundBaseDto 结构体
type RefundBaseDto struct {
	// 退款状态描述
	RefundStatusDesc string `json:"refund_status_desc,omitempty" xml:"refund_status_desc,omitempty"`
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 退款状态
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退款金额/分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}

var poolRefundBaseDto = sync.Pool{
	New: func() any {
		return new(RefundBaseDto)
	},
}

// GetRefundBaseDto() 从对象池中获取RefundBaseDto
func GetRefundBaseDto() *RefundBaseDto {
	return poolRefundBaseDto.Get().(*RefundBaseDto)
}

// ReleaseRefundBaseDto 释放RefundBaseDto
func ReleaseRefundBaseDto(v *RefundBaseDto) {
	v.RefundStatusDesc = ""
	v.BizOrderId = 0
	v.RefundStatus = 0
	v.RefundFee = 0
	poolRefundBaseDto.Put(v)
}
