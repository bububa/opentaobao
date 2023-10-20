package baichuan

import (
	"sync"
)

// PasswordRuleResultDto 结构体
type PasswordRuleResultDto struct {
	// miaoPasswordRegular
	MiaoPasswordRegulars []string `json:"miao_password_regulars,omitempty" xml:"miao_password_regulars>string,omitempty"`
	// passwordRegular
	PasswordRegulars []string `json:"password_regulars,omitempty" xml:"password_regulars>string,omitempty"`
	// level
	Level string `json:"level,omitempty" xml:"level,omitempty"`
}

var poolPasswordRuleResultDto = sync.Pool{
	New: func() any {
		return new(PasswordRuleResultDto)
	},
}

// GetPasswordRuleResultDto() 从对象池中获取PasswordRuleResultDto
func GetPasswordRuleResultDto() *PasswordRuleResultDto {
	return poolPasswordRuleResultDto.Get().(*PasswordRuleResultDto)
}

// ReleasePasswordRuleResultDto 释放PasswordRuleResultDto
func ReleasePasswordRuleResultDto(v *PasswordRuleResultDto) {
	v.MiaoPasswordRegulars = v.MiaoPasswordRegulars[:0]
	v.PasswordRegulars = v.PasswordRegulars[:0]
	v.Level = ""
	poolPasswordRuleResultDto.Put(v)
}
