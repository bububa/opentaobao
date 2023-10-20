package wdk

import (
	"sync"
)

// CsApplySubOrderDto 结构体
type CsApplySubOrderDto struct {
	// 渠道子订单号，淘鲜达渠道为TP子单号
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 申请子单退款金额
	RefundFee float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 申请子单退货数量
	RefundAmount float64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}

var poolCsApplySubOrderDto = sync.Pool{
	New: func() any {
		return new(CsApplySubOrderDto)
	},
}

// GetCsApplySubOrderDto() 从对象池中获取CsApplySubOrderDto
func GetCsApplySubOrderDto() *CsApplySubOrderDto {
	return poolCsApplySubOrderDto.Get().(*CsApplySubOrderDto)
}

// ReleaseCsApplySubOrderDto 释放CsApplySubOrderDto
func ReleaseCsApplySubOrderDto(v *CsApplySubOrderDto) {
	v.OutSubOrderId = ""
	v.RefundFee = 0
	v.RefundAmount = 0
	poolCsApplySubOrderDto.Put(v)
}
