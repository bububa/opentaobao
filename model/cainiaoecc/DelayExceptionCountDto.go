package cainiaoecc

import (
	"sync"
)

// DelayExceptionCountDto 结构体
type DelayExceptionCountDto struct {
	// 扩展字段
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
	// 异常总数
	ExceptionNum int64 `json:"exception_num,omitempty" xml:"exception_num,omitempty"`
	// 商家Id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// CP异常回复总数
	CpReplyExceptionNum int64 `json:"cp_reply_exception_num,omitempty" xml:"cp_reply_exception_num,omitempty"`
}

var poolDelayExceptionCountDto = sync.Pool{
	New: func() any {
		return new(DelayExceptionCountDto)
	},
}

// GetDelayExceptionCountDto() 从对象池中获取DelayExceptionCountDto
func GetDelayExceptionCountDto() *DelayExceptionCountDto {
	return poolDelayExceptionCountDto.Get().(*DelayExceptionCountDto)
}

// ReleaseDelayExceptionCountDto 释放DelayExceptionCountDto
func ReleaseDelayExceptionCountDto(v *DelayExceptionCountDto) {
	v.ExtendFields = ""
	v.ExceptionNum = 0
	v.SellerId = 0
	v.CpReplyExceptionNum = 0
	poolDelayExceptionCountDto.Put(v)
}
