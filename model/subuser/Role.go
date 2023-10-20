package subuser

import (
	"sync"
)

// Role 结构体
type Role struct {
	// 所拥有权限
	Permissions []Permission `json:"permissions,omitempty" xml:"permissions>permission,omitempty"`
	// 角色描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 角色名
	RoleName string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 角色id
	RoleId int64 `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// 卖家Id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolRole = sync.Pool{
	New: func() any {
		return new(Role)
	},
}

// GetRole() 从对象池中获取Role
func GetRole() *Role {
	return poolRole.Get().(*Role)
}

// ReleaseRole 释放Role
func ReleaseRole(v *Role) {
	v.Permissions = v.Permissions[:0]
	v.Description = ""
	v.ModifiedTime = ""
	v.CreateTime = ""
	v.RoleName = ""
	v.RoleId = 0
	v.SellerId = 0
	poolRole.Put(v)
}
