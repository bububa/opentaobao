package alidoc

import (
	"sync"
)

// PrescriptionQueryDto 结构体
type PrescriptionQueryDto struct {
	// 订单号多个以逗号分开
	BizOrderIds string `json:"biz_order_ids,omitempty" xml:"biz_order_ids,omitempty"`
}

var poolPrescriptionQueryDto = sync.Pool{
	New: func() any {
		return new(PrescriptionQueryDto)
	},
}

// GetPrescriptionQueryDto() 从对象池中获取PrescriptionQueryDto
func GetPrescriptionQueryDto() *PrescriptionQueryDto {
	return poolPrescriptionQueryDto.Get().(*PrescriptionQueryDto)
}

// ReleasePrescriptionQueryDto 释放PrescriptionQueryDto
func ReleasePrescriptionQueryDto(v *PrescriptionQueryDto) {
	v.BizOrderIds = ""
	poolPrescriptionQueryDto.Put(v)
}
