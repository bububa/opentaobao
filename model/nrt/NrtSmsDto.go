package nrt

import (
	"sync"
)

// NrtSmsDto 结构体
type NrtSmsDto struct {
	// 短信验证码
	SmsCode string `json:"sms_code,omitempty" xml:"sms_code,omitempty"`
}

var poolNrtSmsDto = sync.Pool{
	New: func() any {
		return new(NrtSmsDto)
	},
}

// GetNrtSmsDto() 从对象池中获取NrtSmsDto
func GetNrtSmsDto() *NrtSmsDto {
	return poolNrtSmsDto.Get().(*NrtSmsDto)
}

// ReleaseNrtSmsDto 释放NrtSmsDto
func ReleaseNrtSmsDto(v *NrtSmsDto) {
	v.SmsCode = ""
	poolNrtSmsDto.Put(v)
}
