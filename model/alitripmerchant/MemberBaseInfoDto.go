package alitripmerchant

import (
	"sync"
)

// MemberBaseInfoDto 结构体
type MemberBaseInfoDto struct {
	// 用户英文姓
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// 用户英文名
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// 称呼
	Civility string `json:"civility,omitempty" xml:"civility,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 城市
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 性别
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 是否接受广告通知
	Subscription bool `json:"subscription,omitempty" xml:"subscription,omitempty"`
}

var poolMemberBaseInfoDto = sync.Pool{
	New: func() any {
		return new(MemberBaseInfoDto)
	},
}

// GetMemberBaseInfoDto() 从对象池中获取MemberBaseInfoDto
func GetMemberBaseInfoDto() *MemberBaseInfoDto {
	return poolMemberBaseInfoDto.Get().(*MemberBaseInfoDto)
}

// ReleaseMemberBaseInfoDto 释放MemberBaseInfoDto
func ReleaseMemberBaseInfoDto(v *MemberBaseInfoDto) {
	v.FirstName = ""
	v.LastName = ""
	v.Civility = ""
	v.Language = ""
	v.Country = ""
	v.Gender = ""
	v.Subscription = false
	poolMemberBaseInfoDto.Put(v)
}
