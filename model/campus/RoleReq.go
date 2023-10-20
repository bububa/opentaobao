package campus

import (
	"sync"
)

// RoleReq 结构体
type RoleReq struct {
	// 角色id
	RoleId string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

var poolRoleReq = sync.Pool{
	New: func() any {
		return new(RoleReq)
	},
}

// GetRoleReq() 从对象池中获取RoleReq
func GetRoleReq() *RoleReq {
	return poolRoleReq.Get().(*RoleReq)
}

// ReleaseRoleReq 释放RoleReq
func ReleaseRoleReq(v *RoleReq) {
	v.RoleId = ""
	poolRoleReq.Put(v)
}
