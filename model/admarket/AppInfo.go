package admarket

import (
	"sync"
)

// AppInfo 结构体
type AppInfo struct {
	// ssp的app名称
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// ssp的app包名
	Pkg string `json:"pkg,omitempty" xml:"pkg,omitempty"`
	// app版本名
	VersionName string `json:"version_name,omitempty" xml:"version_name,omitempty"`
	// app版本号
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
	v.AppName = ""
	v.Pkg = ""
	v.VersionName = ""
	v.VersionCode = 0
	poolAppInfo.Put(v)
}
