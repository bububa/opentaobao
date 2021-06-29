package miniappopen

// MiniappInstanceAppAllVersionsDTO 
type MiniappInstanceAppAllVersionsDTO struct {
    // 小程序信息
    AppInfo   *MiniAppEntityTemplateDTO `json:"app_info,omitempty" xml:"app_info,omitempty"`
    // 数据可能延迟，查不到隔会儿再查。
    AppVersionList   []MiniAppInstanceVersionDTO `json:"app_version_list,omitempty" xml:"app_version_list>mini_app_instance_version_dto,omitempty"`
}
