package alisports

import (
	"sync"
)

// AuthAccountInfoDto 结构体
type AuthAccountInfoDto struct {
	// 头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

var poolAuthAccountInfoDto = sync.Pool{
	New: func() any {
		return new(AuthAccountInfoDto)
	},
}

// GetAuthAccountInfoDto() 从对象池中获取AuthAccountInfoDto
func GetAuthAccountInfoDto() *AuthAccountInfoDto {
	return poolAuthAccountInfoDto.Get().(*AuthAccountInfoDto)
}

// ReleaseAuthAccountInfoDto 释放AuthAccountInfoDto
func ReleaseAuthAccountInfoDto(v *AuthAccountInfoDto) {
	v.Avatar = ""
	v.Nick = ""
	v.OpenId = ""
	poolAuthAccountInfoDto.Put(v)
}
