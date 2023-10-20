package moziacl

import (
	"sync"
)

// PermissionPackageEntity 结构体
type PermissionPackageEntity struct {
	// 注册的租户列表
	RegistRealmList []RealmEntity `json:"regist_realm_list,omitempty" xml:"regist_realm_list>realm_entity,omitempty"`
	// 权限套餐code
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 权限套餐中文名
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 权限套餐英文名
	NameEN string `json:"name_e_n,omitempty" xml:"name_e_n,omitempty"`
	// 权限套餐所属应用名
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 权限套餐中文描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 权限套餐英文描述
	DescriptionEN string `json:"description_e_n,omitempty" xml:"description_e_n,omitempty"`
	// 扩展字段
	ExtStr string `json:"ext_str,omitempty" xml:"ext_str,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 包含功能权限数量
	PermissionCount int64 `json:"permission_count,omitempty" xml:"permission_count,omitempty"`
	// 包含角色数量
	RoleCount int64 `json:"role_count,omitempty" xml:"role_count,omitempty"`
	// 包含数据权限数量
	DataPermissionCount int64 `json:"data_permission_count,omitempty" xml:"data_permission_count,omitempty"`
	// 权限套餐创建人
	Creator *BucUser `json:"creator,omitempty" xml:"creator,omitempty"`
}

var poolPermissionPackageEntity = sync.Pool{
	New: func() any {
		return new(PermissionPackageEntity)
	},
}

// GetPermissionPackageEntity() 从对象池中获取PermissionPackageEntity
func GetPermissionPackageEntity() *PermissionPackageEntity {
	return poolPermissionPackageEntity.Get().(*PermissionPackageEntity)
}

// ReleasePermissionPackageEntity 释放PermissionPackageEntity
func ReleasePermissionPackageEntity(v *PermissionPackageEntity) {
	v.RegistRealmList = v.RegistRealmList[:0]
	v.Name = ""
	v.Title = ""
	v.NameEN = ""
	v.AppName = ""
	v.Description = ""
	v.DescriptionEN = ""
	v.ExtStr = ""
	v.CreateTime = ""
	v.PermissionCount = 0
	v.RoleCount = 0
	v.DataPermissionCount = 0
	v.Creator = nil
	poolPermissionPackageEntity.Put(v)
}
