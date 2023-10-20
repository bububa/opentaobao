package mos

import (
	"sync"
)

// CancelBillDto 结构体
type CancelBillDto struct {
	// 结算行集合
	SettleLineNos []string `json:"settle_line_nos,omitempty" xml:"settle_line_nos>string,omitempty"`
	// 取消备注
	CancelComments string `json:"cancel_comments,omitempty" xml:"cancel_comments,omitempty"`
	// 取消类型
	CancelType string `json:"cancel_type,omitempty" xml:"cancel_type,omitempty"`
	// 付款单号
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 备注
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}

var poolCancelBillDto = sync.Pool{
	New: func() any {
		return new(CancelBillDto)
	},
}

// GetCancelBillDto() 从对象池中获取CancelBillDto
func GetCancelBillDto() *CancelBillDto {
	return poolCancelBillDto.Get().(*CancelBillDto)
}

// ReleaseCancelBillDto 释放CancelBillDto
func ReleaseCancelBillDto(v *CancelBillDto) {
	v.SettleLineNos = v.SettleLineNos[:0]
	v.CancelComments = ""
	v.CancelType = ""
	v.BillNo = ""
	v.ExtendParams = ""
	poolCancelBillDto.Put(v)
}
