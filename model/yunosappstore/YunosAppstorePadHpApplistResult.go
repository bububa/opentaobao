package yunosappstore

import (
	"sync"
)

// YunosAppstorePadHpApplistResult 结构体
type YunosAppstorePadHpApplistResult struct {
	// icon
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// packageName
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// developerName
	DeveloperName string `json:"developer_name,omitempty" xml:"developer_name,omitempty"`
	// uri
	Uri string `json:"uri,omitempty" xml:"uri,omitempty"`
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// intro
	Intro string `json:"intro,omitempty" xml:"intro,omitempty"`
	// versionCode
	VersionCode int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
}

var poolYunosAppstorePadHpApplistResult = sync.Pool{
	New: func() any {
		return new(YunosAppstorePadHpApplistResult)
	},
}

// GetYunosAppstorePadHpApplistResult() 从对象池中获取YunosAppstorePadHpApplistResult
func GetYunosAppstorePadHpApplistResult() *YunosAppstorePadHpApplistResult {
	return poolYunosAppstorePadHpApplistResult.Get().(*YunosAppstorePadHpApplistResult)
}

// ReleaseYunosAppstorePadHpApplistResult 释放YunosAppstorePadHpApplistResult
func ReleaseYunosAppstorePadHpApplistResult(v *YunosAppstorePadHpApplistResult) {
	v.Icon = ""
	v.GmtCreate = ""
	v.PackageName = ""
	v.Name = ""
	v.DeveloperName = ""
	v.Uri = ""
	v.Url = ""
	v.Intro = ""
	v.VersionCode = 0
	poolYunosAppstorePadHpApplistResult.Put(v)
}
