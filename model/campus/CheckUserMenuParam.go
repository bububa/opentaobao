package campus

import (
	"sync"
)

// CheckUserMenuParam 结构体
type CheckUserMenuParam struct {
	// 菜单url
	MenuUrl string `json:"menu_url,omitempty" xml:"menu_url,omitempty"`
	// 用户账号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolCheckUserMenuParam = sync.Pool{
	New: func() any {
		return new(CheckUserMenuParam)
	},
}

// GetCheckUserMenuParam() 从对象池中获取CheckUserMenuParam
func GetCheckUserMenuParam() *CheckUserMenuParam {
	return poolCheckUserMenuParam.Get().(*CheckUserMenuParam)
}

// ReleaseCheckUserMenuParam 释放CheckUserMenuParam
func ReleaseCheckUserMenuParam(v *CheckUserMenuParam) {
	v.MenuUrl = ""
	v.UserId = ""
	poolCheckUserMenuParam.Put(v)
}
