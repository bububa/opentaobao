package cloudgame

import (
	"sync"
)

// LoginSessionDto 结构体
type LoginSessionDto struct {
	// user的acesstoken
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// user在mp上的accountid
	MpAccountId int64 `json:"mp_account_id,omitempty" xml:"mp_account_id,omitempty"`
}

var poolLoginSessionDto = sync.Pool{
	New: func() any {
		return new(LoginSessionDto)
	},
}

// GetLoginSessionDto() 从对象池中获取LoginSessionDto
func GetLoginSessionDto() *LoginSessionDto {
	return poolLoginSessionDto.Get().(*LoginSessionDto)
}

// ReleaseLoginSessionDto 释放LoginSessionDto
func ReleaseLoginSessionDto(v *LoginSessionDto) {
	v.AccessToken = ""
	v.MpAccountId = 0
	poolLoginSessionDto.Put(v)
}
