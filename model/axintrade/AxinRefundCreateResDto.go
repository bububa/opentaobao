package axintrade

import (
	"sync"
)

// AxinRefundCreateResDto 结构体
type AxinRefundCreateResDto struct {
	// 请求版本号，用于幂等校验
	ReqVersion string `json:"req_version,omitempty" xml:"req_version,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 实退金额
	ActualFee int64 `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}

var poolAxinRefundCreateResDto = sync.Pool{
	New: func() any {
		return new(AxinRefundCreateResDto)
	},
}

// GetAxinRefundCreateResDto() 从对象池中获取AxinRefundCreateResDto
func GetAxinRefundCreateResDto() *AxinRefundCreateResDto {
	return poolAxinRefundCreateResDto.Get().(*AxinRefundCreateResDto)
}

// ReleaseAxinRefundCreateResDto 释放AxinRefundCreateResDto
func ReleaseAxinRefundCreateResDto(v *AxinRefundCreateResDto) {
	v.ReqVersion = ""
	v.OuterOrderId = ""
	v.ActualFee = 0
	v.RefundFee = 0
	poolAxinRefundCreateResDto.Put(v)
}
