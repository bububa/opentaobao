package crm

import (
	"sync"
)

// MemberAccountDto 结构体
type MemberAccountDto struct {
	// 等级名称
	GradeName string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 等级编号
	Grade int64 `json:"grade,omitempty" xml:"grade,omitempty"`
	// bindStatus 1：绑卡（已经是线下会员线上未绑定，或者解绑后再绑定），2：注册
	BindStatus int64 `json:"bind_status,omitempty" xml:"bind_status,omitempty"`
}

var poolMemberAccountDto = sync.Pool{
	New: func() any {
		return new(MemberAccountDto)
	},
}

// GetMemberAccountDto() 从对象池中获取MemberAccountDto
func GetMemberAccountDto() *MemberAccountDto {
	return poolMemberAccountDto.Get().(*MemberAccountDto)
}

// ReleaseMemberAccountDto 释放MemberAccountDto
func ReleaseMemberAccountDto(v *MemberAccountDto) {
	v.GradeName = ""
	v.GmtCreate = ""
	v.Grade = 0
	v.BindStatus = 0
	poolMemberAccountDto.Put(v)
}
