package user

import (
	"sync"
)

// OpenUserInfoDto 结构体
type OpenUserInfoDto struct {
	// 混淆字符串
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 头像链接
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// snsNick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
}

var poolOpenUserInfoDto = sync.Pool{
	New: func() any {
		return new(OpenUserInfoDto)
	},
}

// GetOpenUserInfoDto() 从对象池中获取OpenUserInfoDto
func GetOpenUserInfoDto() *OpenUserInfoDto {
	return poolOpenUserInfoDto.Get().(*OpenUserInfoDto)
}

// ReleaseOpenUserInfoDto 释放OpenUserInfoDto
func ReleaseOpenUserInfoDto(v *OpenUserInfoDto) {
	v.OpenId = ""
	v.Avatar = ""
	v.Nick = ""
	poolOpenUserInfoDto.Put(v)
}
