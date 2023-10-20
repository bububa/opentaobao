package idle

import (
	"sync"
)

// IdleGooseFishUserInfoVo 结构体
type IdleGooseFishUserInfoVo struct {
	// 闲鱼头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 用户闲鱼昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 0男生 1女生
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
}

var poolIdleGooseFishUserInfoVo = sync.Pool{
	New: func() any {
		return new(IdleGooseFishUserInfoVo)
	},
}

// GetIdleGooseFishUserInfoVo() 从对象池中获取IdleGooseFishUserInfoVo
func GetIdleGooseFishUserInfoVo() *IdleGooseFishUserInfoVo {
	return poolIdleGooseFishUserInfoVo.Get().(*IdleGooseFishUserInfoVo)
}

// ReleaseIdleGooseFishUserInfoVo 释放IdleGooseFishUserInfoVo
func ReleaseIdleGooseFishUserInfoVo(v *IdleGooseFishUserInfoVo) {
	v.Avatar = ""
	v.NickName = ""
	v.Gender = 0
	poolIdleGooseFishUserInfoVo.Put(v)
}
