package jst

import (
	"sync"
)

// QuerySmsSignDto 结构体
type QuerySmsSignDto struct {
	// 被拒绝的原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 申请的签名
	SignName string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
	// 签名创建时间
	CreateDate string `json:"create_date,omitempty" xml:"create_date,omitempty"`
	// 拒绝
	SignStatus int64 `json:"sign_status,omitempty" xml:"sign_status,omitempty"`
}

var poolQuerySmsSignDto = sync.Pool{
	New: func() any {
		return new(QuerySmsSignDto)
	},
}

// GetQuerySmsSignDto() 从对象池中获取QuerySmsSignDto
func GetQuerySmsSignDto() *QuerySmsSignDto {
	return poolQuerySmsSignDto.Get().(*QuerySmsSignDto)
}

// ReleaseQuerySmsSignDto 释放QuerySmsSignDto
func ReleaseQuerySmsSignDto(v *QuerySmsSignDto) {
	v.Reason = ""
	v.SignName = ""
	v.CreateDate = ""
	v.SignStatus = 0
	poolQuerySmsSignDto.Put(v)
}
