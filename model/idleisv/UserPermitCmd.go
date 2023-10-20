package idleisv

import (
	"sync"
)

// UserPermitCmd 结构体
type UserPermitCmd struct {
	// 当前用户的所属业务类型编码，优品&amp;开放平台业务 默认使用 IDLE_TOP
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}

var poolUserPermitCmd = sync.Pool{
	New: func() any {
		return new(UserPermitCmd)
	},
}

// GetUserPermitCmd() 从对象池中获取UserPermitCmd
func GetUserPermitCmd() *UserPermitCmd {
	return poolUserPermitCmd.Get().(*UserPermitCmd)
}

// ReleaseUserPermitCmd 释放UserPermitCmd
func ReleaseUserPermitCmd(v *UserPermitCmd) {
	v.BizCode = ""
	poolUserPermitCmd.Put(v)
}
