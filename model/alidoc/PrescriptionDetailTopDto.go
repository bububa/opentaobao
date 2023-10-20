package alidoc

import (
	"sync"
)

// PrescriptionDetailTopDto 结构体
type PrescriptionDetailTopDto struct {
	// 订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 处方图地址
	PrescriptionPicUrl string `json:"prescription_pic_url,omitempty" xml:"prescription_pic_url,omitempty"`
}

var poolPrescriptionDetailTopDto = sync.Pool{
	New: func() any {
		return new(PrescriptionDetailTopDto)
	},
}

// GetPrescriptionDetailTopDto() 从对象池中获取PrescriptionDetailTopDto
func GetPrescriptionDetailTopDto() *PrescriptionDetailTopDto {
	return poolPrescriptionDetailTopDto.Get().(*PrescriptionDetailTopDto)
}

// ReleasePrescriptionDetailTopDto 释放PrescriptionDetailTopDto
func ReleasePrescriptionDetailTopDto(v *PrescriptionDetailTopDto) {
	v.BizOrderId = ""
	v.PrescriptionPicUrl = ""
	poolPrescriptionDetailTopDto.Put(v)
}
