package campus

import (
	"sync"
)

// UserRoleQueryParam 结构体
type UserRoleQueryParam struct {
	// 多应用
	AppKeys []string `json:"app_keys,omitempty" xml:"app_keys>string,omitempty"`
	// 用户账号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 是否支持有效期过滤
	ReturnNotEffective bool `json:"return_not_effective,omitempty" xml:"return_not_effective,omitempty"`
}

var poolUserRoleQueryParam = sync.Pool{
	New: func() any {
		return new(UserRoleQueryParam)
	},
}

// GetUserRoleQueryParam() 从对象池中获取UserRoleQueryParam
func GetUserRoleQueryParam() *UserRoleQueryParam {
	return poolUserRoleQueryParam.Get().(*UserRoleQueryParam)
}

// ReleaseUserRoleQueryParam 释放UserRoleQueryParam
func ReleaseUserRoleQueryParam(v *UserRoleQueryParam) {
	v.AppKeys = v.AppKeys[:0]
	v.UserId = ""
	v.ReturnNotEffective = false
	poolUserRoleQueryParam.Put(v)
}
