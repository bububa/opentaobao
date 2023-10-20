package cainiaoecc

import (
	"sync"
)

// DelayExceptionDetailDto 结构体
type DelayExceptionDetailDto struct {
	// CP回复列表
	CpReplyList []string `json:"cp_reply_list,omitempty" xml:"cp_reply_list>string,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 异常类型code
	ExceptionCode string `json:"exception_code,omitempty" xml:"exception_code,omitempty"`
	// 异常类型名称
	ExceptionName string `json:"exception_name,omitempty" xml:"exception_name,omitempty"`
	// 商家Id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolDelayExceptionDetailDto = sync.Pool{
	New: func() any {
		return new(DelayExceptionDetailDto)
	},
}

// GetDelayExceptionDetailDto() 从对象池中获取DelayExceptionDetailDto
func GetDelayExceptionDetailDto() *DelayExceptionDetailDto {
	return poolDelayExceptionDetailDto.Get().(*DelayExceptionDetailDto)
}

// ReleaseDelayExceptionDetailDto 释放DelayExceptionDetailDto
func ReleaseDelayExceptionDetailDto(v *DelayExceptionDetailDto) {
	v.CpReplyList = v.CpReplyList[:0]
	v.MailNo = ""
	v.ExceptionCode = ""
	v.ExceptionName = ""
	v.SellerId = 0
	poolDelayExceptionDetailDto.Put(v)
}
