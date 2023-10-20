package admarket

import (
	"sync"
)

// SdkInfo 结构体
type SdkInfo struct {
	// sdk版本名
	VersionName string `json:"version_name,omitempty" xml:"version_name,omitempty"`
	// sdk版本号
	VersionCode int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
}

var poolSdkInfo = sync.Pool{
	New: func() any {
		return new(SdkInfo)
	},
}

// GetSdkInfo() 从对象池中获取SdkInfo
func GetSdkInfo() *SdkInfo {
	return poolSdkInfo.Get().(*SdkInfo)
}

// ReleaseSdkInfo 释放SdkInfo
func ReleaseSdkInfo(v *SdkInfo) {
	v.VersionName = ""
	v.VersionCode = 0
	poolSdkInfo.Put(v)
}
