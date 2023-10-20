package miniappopen

import (
	"sync"
)

// MiniAppInstanceVersionDto 结构体
type MiniAppInstanceVersionDto struct {
	// 小程序版本号
	AppVersion string `json:"app_version,omitempty" xml:"app_version,omitempty"`
	// 发布端
	Client string `json:"client,omitempty" xml:"client,omitempty"`
	// 小程序app_id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 版本状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 模板id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 模板版本
	TemplateVersion string `json:"template_version,omitempty" xml:"template_version,omitempty"`
	// 扩展信息
	ExtJson string `json:"ext_json,omitempty" xml:"ext_json,omitempty"`
	// 版本链接。上线状态为线上地址，预览状态为预览地址，下线状态为空。
	AppUrl string `json:"app_url,omitempty" xml:"app_url,omitempty"`
	// 实例版本
	Version string `json:"version,omitempty" xml:"version,omitempty"`
}

var poolMiniAppInstanceVersionDto = sync.Pool{
	New: func() any {
		return new(MiniAppInstanceVersionDto)
	},
}

// GetMiniAppInstanceVersionDto() 从对象池中获取MiniAppInstanceVersionDto
func GetMiniAppInstanceVersionDto() *MiniAppInstanceVersionDto {
	return poolMiniAppInstanceVersionDto.Get().(*MiniAppInstanceVersionDto)
}

// ReleaseMiniAppInstanceVersionDto 释放MiniAppInstanceVersionDto
func ReleaseMiniAppInstanceVersionDto(v *MiniAppInstanceVersionDto) {
	v.AppVersion = ""
	v.Client = ""
	v.AppId = ""
	v.Status = ""
	v.TemplateId = ""
	v.TemplateVersion = ""
	v.ExtJson = ""
	v.AppUrl = ""
	v.Version = ""
	poolMiniAppInstanceVersionDto.Put(v)
}
