package miniappopen

import (
	"sync"
)

// MiniappInstanceAppAllVersionsDto 结构体
type MiniappInstanceAppAllVersionsDto struct {
	// 数据可能延迟，查不到隔会儿再查。
	AppVersionList []MiniAppInstanceVersionDto `json:"app_version_list,omitempty" xml:"app_version_list>mini_app_instance_version_dto,omitempty"`
	// 小程序信息
	AppInfo *MiniAppEntityTemplateDto `json:"app_info,omitempty" xml:"app_info,omitempty"`
	// 实体信息
	MiniAppEntityTemplateDto *MiniAppEntityTemplateDto `json:"mini_app_entity_template_dto,omitempty" xml:"mini_app_entity_template_dto,omitempty"`
}

var poolMiniappInstanceAppAllVersionsDto = sync.Pool{
	New: func() any {
		return new(MiniappInstanceAppAllVersionsDto)
	},
}

// GetMiniappInstanceAppAllVersionsDto() 从对象池中获取MiniappInstanceAppAllVersionsDto
func GetMiniappInstanceAppAllVersionsDto() *MiniappInstanceAppAllVersionsDto {
	return poolMiniappInstanceAppAllVersionsDto.Get().(*MiniappInstanceAppAllVersionsDto)
}

// ReleaseMiniappInstanceAppAllVersionsDto 释放MiniappInstanceAppAllVersionsDto
func ReleaseMiniappInstanceAppAllVersionsDto(v *MiniappInstanceAppAllVersionsDto) {
	v.AppVersionList = v.AppVersionList[:0]
	v.AppInfo = nil
	v.MiniAppEntityTemplateDto = nil
	poolMiniappInstanceAppAllVersionsDto.Put(v)
}
