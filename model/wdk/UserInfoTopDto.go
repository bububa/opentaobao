package wdk

import (
	"sync"
)

// UserInfoTopDto 结构体
type UserInfoTopDto struct {
	// 人员account
	UserAccount string `json:"user_account,omitempty" xml:"user_account,omitempty"`
	// 人员名字
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

var poolUserInfoTopDto = sync.Pool{
	New: func() any {
		return new(UserInfoTopDto)
	},
}

// GetUserInfoTopDto() 从对象池中获取UserInfoTopDto
func GetUserInfoTopDto() *UserInfoTopDto {
	return poolUserInfoTopDto.Get().(*UserInfoTopDto)
}

// ReleaseUserInfoTopDto 释放UserInfoTopDto
func ReleaseUserInfoTopDto(v *UserInfoTopDto) {
	v.UserAccount = ""
	v.UserName = ""
	poolUserInfoTopDto.Put(v)
}
