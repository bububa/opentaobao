package tmallcar

import (
	"sync"
)

// CarSubPayOrderDto 结构体
type CarSubPayOrderDto struct {
	// 阶段名称
	StepName string `json:"step_name,omitempty" xml:"step_name,omitempty"`
	// 阶段号
	StepNo int64 `json:"step_no,omitempty" xml:"step_no,omitempty"`
	// 需要付款的金额，单位分
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 实际付款金额，单位分
	ActualPayFee int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 子支付单id
	SubPayOrderId int64 `json:"sub_pay_order_id,omitempty" xml:"sub_pay_order_id,omitempty"`
}

var poolCarSubPayOrderDto = sync.Pool{
	New: func() any {
		return new(CarSubPayOrderDto)
	},
}

// GetCarSubPayOrderDto() 从对象池中获取CarSubPayOrderDto
func GetCarSubPayOrderDto() *CarSubPayOrderDto {
	return poolCarSubPayOrderDto.Get().(*CarSubPayOrderDto)
}

// ReleaseCarSubPayOrderDto 释放CarSubPayOrderDto
func ReleaseCarSubPayOrderDto(v *CarSubPayOrderDto) {
	v.StepName = ""
	v.StepNo = 0
	v.PayFee = 0
	v.ActualPayFee = 0
	v.PayStatus = 0
	v.SubPayOrderId = 0
	poolCarSubPayOrderDto.Put(v)
}
