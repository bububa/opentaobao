package axintrade

import (
	"sync"
)

// AxinRefundCreateDto 结构体
type AxinRefundCreateDto struct {
	// 扩展字段(k:v结构)
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 退款说明
	RefundExplanation string `json:"refund_explanation,omitempty" xml:"refund_explanation,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 请求版本号，用于幂等校验
	ReqVersion string `json:"req_version,omitempty" xml:"req_version,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}

var poolAxinRefundCreateDto = sync.Pool{
	New: func() any {
		return new(AxinRefundCreateDto)
	},
}

// GetAxinRefundCreateDto() 从对象池中获取AxinRefundCreateDto
func GetAxinRefundCreateDto() *AxinRefundCreateDto {
	return poolAxinRefundCreateDto.Get().(*AxinRefundCreateDto)
}

// ReleaseAxinRefundCreateDto 释放AxinRefundCreateDto
func ReleaseAxinRefundCreateDto(v *AxinRefundCreateDto) {
	v.Attributes = ""
	v.RefundExplanation = ""
	v.RefundReason = ""
	v.ReqVersion = ""
	v.OuterOrderId = ""
	v.RefundFee = 0
	poolAxinRefundCreateDto.Put(v)
}
