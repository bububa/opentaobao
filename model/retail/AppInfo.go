package retail

import (
	"sync"
)

// AppInfo 结构体
type AppInfo struct {
	// 系统自动生成
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
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
	poolAppInfo.Put(v)
}
