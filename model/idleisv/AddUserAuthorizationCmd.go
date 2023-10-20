package idleisv

import (
	"sync"
)

// AddUserAuthorizationCmd 结构体
type AddUserAuthorizationCmd struct {
	// 需要为用户添加的权限类型，具体类型可以参考api描述的语雀文档
	AuthType string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
}

var poolAddUserAuthorizationCmd = sync.Pool{
	New: func() any {
		return new(AddUserAuthorizationCmd)
	},
}

// GetAddUserAuthorizationCmd() 从对象池中获取AddUserAuthorizationCmd
func GetAddUserAuthorizationCmd() *AddUserAuthorizationCmd {
	return poolAddUserAuthorizationCmd.Get().(*AddUserAuthorizationCmd)
}

// ReleaseAddUserAuthorizationCmd 释放AddUserAuthorizationCmd
func ReleaseAddUserAuthorizationCmd(v *AddUserAuthorizationCmd) {
	v.AuthType = ""
	poolAddUserAuthorizationCmd.Put(v)
}
