package alihealth2

// SourcesReturnVo 
type SourcesReturnVo struct {
    // 地区编码
    AreaCode   string `json:"area_code,omitempty" xml:"area_code,omitempty"`
    // 地区
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 医院id
    HosId   string `json:"hos_id,omitempty" xml:"hos_id,omitempty"`
    // 医院名称
    HosName   string `json:"hos_name,omitempty" xml:"hos_name,omitempty"`
    // 是否有号 0 无号 1有号
    SourceStatus   string `json:"source_status,omitempty" xml:"source_status,omitempty"`
}
