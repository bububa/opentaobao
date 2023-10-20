package mos

import (
	"sync"
)

// CallDispatcherDto 结构体
type CallDispatcherDto struct {
	// 包裹信息
	CodeInfoList []CodeInfoDto `json:"code_info_list,omitempty" xml:"code_info_list>code_info_dto,omitempty"`
	// 主单号
	ParentNo string `json:"parent_no,omitempty" xml:"parent_no,omitempty"`
	// 收货信息
	ReceiveInfo *DeliveryCustomDto `json:"receive_info,omitempty" xml:"receive_info,omitempty"`
}

var poolCallDispatcherDto = sync.Pool{
	New: func() any {
		return new(CallDispatcherDto)
	},
}

// GetCallDispatcherDto() 从对象池中获取CallDispatcherDto
func GetCallDispatcherDto() *CallDispatcherDto {
	return poolCallDispatcherDto.Get().(*CallDispatcherDto)
}

// ReleaseCallDispatcherDto 释放CallDispatcherDto
func ReleaseCallDispatcherDto(v *CallDispatcherDto) {
	v.CodeInfoList = v.CodeInfoList[:0]
	v.ParentNo = ""
	v.ReceiveInfo = nil
	poolCallDispatcherDto.Put(v)
}
