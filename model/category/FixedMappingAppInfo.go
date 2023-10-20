package category

import (
	"sync"
)

// FixedMappingAppInfo 结构体
type FixedMappingAppInfo struct {
	// 访问映射接口密码
	Password string `json:"password,omitempty" xml:"password,omitempty"`
	// 访问映射接口账号
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
}

var poolFixedMappingAppInfo = sync.Pool{
	New: func() any {
		return new(FixedMappingAppInfo)
	},
}

// GetFixedMappingAppInfo() 从对象池中获取FixedMappingAppInfo
func GetFixedMappingAppInfo() *FixedMappingAppInfo {
	return poolFixedMappingAppInfo.Get().(*FixedMappingAppInfo)
}

// ReleaseFixedMappingAppInfo 释放FixedMappingAppInfo
func ReleaseFixedMappingAppInfo(v *FixedMappingAppInfo) {
	v.Password = ""
	v.AppName = ""
	poolFixedMappingAppInfo.Put(v)
}
