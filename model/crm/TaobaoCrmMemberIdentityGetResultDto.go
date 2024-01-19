package crm

import (
	"sync"
)

// TaobaoCrmMemberIdentityGetResultDto 结构体
type TaobaoCrmMemberIdentityGetResultDto struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// total
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// result
	MemberInfo *MemberAccountDto `json:"member_info,omitempty" xml:"member_info,omitempty"`
}

var poolTaobaoCrmMemberIdentityGetResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMemberIdentityGetResultDto)
	},
}

// GetTaobaoCrmMemberIdentityGetResultDto() 从对象池中获取TaobaoCrmMemberIdentityGetResultDto
func GetTaobaoCrmMemberIdentityGetResultDto() *TaobaoCrmMemberIdentityGetResultDto {
	return poolTaobaoCrmMemberIdentityGetResultDto.Get().(*TaobaoCrmMemberIdentityGetResultDto)
}

// ReleaseTaobaoCrmMemberIdentityGetResultDto 释放TaobaoCrmMemberIdentityGetResultDto
func ReleaseTaobaoCrmMemberIdentityGetResultDto(v *TaobaoCrmMemberIdentityGetResultDto) {
	v.Code = ""
	v.Msg = ""
	v.Total = 0
	v.MemberInfo = nil
	poolTaobaoCrmMemberIdentityGetResultDto.Put(v)
}
