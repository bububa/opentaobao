package wdk

import (
	"sync"
)

// MemberInfoDto 结构体
type MemberInfoDto struct {
	// 淘宝用户昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 对应淘宝账号的OpenUID
	UicId string `json:"uic_id,omitempty" xml:"uic_id,omitempty"`
}

var poolMemberInfoDto = sync.Pool{
	New: func() any {
		return new(MemberInfoDto)
	},
}

// GetMemberInfoDto() 从对象池中获取MemberInfoDto
func GetMemberInfoDto() *MemberInfoDto {
	return poolMemberInfoDto.Get().(*MemberInfoDto)
}

// ReleaseMemberInfoDto 释放MemberInfoDto
func ReleaseMemberInfoDto(v *MemberInfoDto) {
	v.Nick = ""
	v.UicId = ""
	poolMemberInfoDto.Put(v)
}
