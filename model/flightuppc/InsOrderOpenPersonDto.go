package flightuppc

import (
	"sync"
)

// InsOrderOpenPersonDto 结构体
type InsOrderOpenPersonDto struct {
	// idCardNo(md5脱敏)
	IdCardNo string `json:"id_card_no,omitempty" xml:"id_card_no,omitempty"`
	// 子保单号
	PolicyNo string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// 保险订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// idCardType
	IdCardType int64 `json:"id_card_type,omitempty" xml:"id_card_type,omitempty"`
	// 外部订单号
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

var poolInsOrderOpenPersonDto = sync.Pool{
	New: func() any {
		return new(InsOrderOpenPersonDto)
	},
}

// GetInsOrderOpenPersonDto() 从对象池中获取InsOrderOpenPersonDto
func GetInsOrderOpenPersonDto() *InsOrderOpenPersonDto {
	return poolInsOrderOpenPersonDto.Get().(*InsOrderOpenPersonDto)
}

// ReleaseInsOrderOpenPersonDto 释放InsOrderOpenPersonDto
func ReleaseInsOrderOpenPersonDto(v *InsOrderOpenPersonDto) {
	v.IdCardNo = ""
	v.PolicyNo = ""
	v.TcOrderId = 0
	v.IdCardType = 0
	v.OutOrderId = 0
	poolInsOrderOpenPersonDto.Put(v)
}
