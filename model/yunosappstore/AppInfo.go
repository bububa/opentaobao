package yunosappstore

import (
	"sync"
)

// AppInfo 结构体
type AppInfo struct {
	// 应用icon
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用包名
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 应用名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用版本名称
	VersionName string `json:"version_name,omitempty" xml:"version_name,omitempty"`
	// 应用版本号
	VersionCode int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
}

var poolAppInfo = sync.Pool{
	New: func() any {
		return new(AppInfo)
	},
}

// GetAppInfo() 从对象池中获取AppInfo
func GetAppInfo() *AppInfo {
	return poolAppInfo.Get().(*AppInfo)
}

// ReleaseAppInfo 释放AppInfo
func ReleaseAppInfo(v *AppInfo) {
	v.Icon = ""
	v.PackageName = ""
	v.Name = ""
	v.VersionName = ""
	v.VersionCode = 0
	poolAppInfo.Put(v)
}
