package util

import (
	"sync"
)

// Dto 结构体
type Dto struct {
	// 自定义配置信息
	ConfigJson string `json:"config_json,omitempty" xml:"config_json,omitempty"`
	// 资源文件下载地址
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	// 二维码地址
	Qrcode string `json:"qrcode,omitempty" xml:"qrcode,omitempty"`
	// 资源描述
	UpdateDesc string `json:"update_desc,omitempty" xml:"update_desc,omitempty"`
	// 资源版本名
	VersionName string `json:"version_name,omitempty" xml:"version_name,omitempty"`
	// 图标地址
	IconUrl string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// 资源MD5
	Md5 string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 应用名
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 主包名
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 兼容主包名
	OldPackageName string `json:"old_package_name,omitempty" xml:"old_package_name,omitempty"`
	// 自定义标识符
	Identifier string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 资源状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 资源大小
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 资源描述
	VersionCode int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
	// 是否自动安装
	AutoInstall bool `json:"auto_install,omitempty" xml:"auto_install,omitempty"`
}

var poolDto = sync.Pool{
	New: func() any {
		return new(Dto)
	},
}

// GetDto() 从对象池中获取Dto
func GetDto() *Dto {
	return poolDto.Get().(*Dto)
}

// ReleaseDto 释放Dto
func ReleaseDto(v *Dto) {
	v.ConfigJson = ""
	v.DownloadUrl = ""
	v.Qrcode = ""
	v.UpdateDesc = ""
	v.VersionName = ""
	v.IconUrl = ""
	v.Md5 = ""
	v.AppName = ""
	v.PackageName = ""
	v.OldPackageName = ""
	v.Identifier = ""
	v.Status = 0
	v.Size = 0
	v.VersionCode = 0
	v.AutoInstall = false
	poolDto.Put(v)
}
