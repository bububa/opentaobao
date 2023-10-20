package ott

import (
	"sync"
)

// MemberDto 结构体
type MemberDto struct {
	// characterName
	CharacterName string `json:"character_name,omitempty" xml:"character_name,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// personId
	PersonId string `json:"person_id,omitempty" xml:"person_id,omitempty"`
	// role
	Role string `json:"role,omitempty" xml:"role,omitempty"`
}

var poolMemberDto = sync.Pool{
	New: func() any {
		return new(MemberDto)
	},
}

// GetMemberDto() 从对象池中获取MemberDto
func GetMemberDto() *MemberDto {
	return poolMemberDto.Get().(*MemberDto)
}

// ReleaseMemberDto 释放MemberDto
func ReleaseMemberDto(v *MemberDto) {
	v.CharacterName = ""
	v.Name = ""
	v.PersonId = ""
	v.Role = ""
	poolMemberDto.Put(v)
}
