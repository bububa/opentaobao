package miniappopen

import (
	"sync"
)

// MiniAppEntityDto 结构体
type MiniAppEntityDto struct {
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 小程序appId
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 小程序名称
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 小程序描述
	AppDescription string `json:"app_description,omitempty" xml:"app_description,omitempty"`
	// icon
	AppIcon string `json:"app_icon,omitempty" xml:"app_icon,omitempty"`
	// 当前版本
	OnlineVersion string `json:"online_version,omitempty" xml:"online_version,omitempty"`
	// 线上码
	OnlineCode string `json:"online_code,omitempty" xml:"online_code,omitempty"`
}

var poolMiniAppEntityDto = sync.Pool{
	New: func() any {
		return new(MiniAppEntityDto)
	},
}

// GetMiniAppEntityDto() 从对象池中获取MiniAppEntityDto
func GetMiniAppEntityDto() *MiniAppEntityDto {
	return poolMiniAppEntityDto.Get().(*MiniAppEntityDto)
}

// ReleaseMiniAppEntityDto 释放MiniAppEntityDto
func ReleaseMiniAppEntityDto(v *MiniAppEntityDto) {
	v.Appkey = ""
	v.Id = ""
	v.AppName = ""
	v.AppDescription = ""
	v.AppIcon = ""
	v.OnlineVersion = ""
	v.OnlineCode = ""
	poolMiniAppEntityDto.Put(v)
}
