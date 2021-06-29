package damai

// ThirdProjectPushOpenParam 
type ThirdProjectPushOpenParam struct {
    // 城市id
    CityId   int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
    // 图片url
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    // 项目id
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    // 项目名称
    ProjectName   string `json:"project_name,omitempty" xml:"project_name,omitempty"`
    // 推送时间
    PushTime   string `json:"push_time,omitempty" xml:"push_time,omitempty"`
    // 商户密钥
    SupplierSecret   string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
    // 系统id
    SystemId   int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
    // 场馆id
    VenueId   int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
}
