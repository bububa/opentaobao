package alihouse

// BaseMetroLineDto 
type BaseMetroLineDto struct {
    // 所属商圈
    BusinessDistrict   string `json:"business_district,omitempty" xml:"business_district,omitempty"`
    // 站点编号
    SiteCode   string `json:"site_code,omitempty" xml:"site_code,omitempty"`
    // 线路编号
    LineCode   string `json:"line_code,omitempty" xml:"line_code,omitempty"`
    // 站所属区域
    AreaId   int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
    // 城市代码
    CityId   int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
    // 地铁站所属线路
    MetroLine   string `json:"metro_line,omitempty" xml:"metro_line,omitempty"`
    // 地铁站名称
    MetroName   string `json:"metro_name,omitempty" xml:"metro_name,omitempty"`
    // 地铁站代码
    MetroCode   string `json:"metro_code,omitempty" xml:"metro_code,omitempty"`
    // 外部ID -- 唯一
    OuterMetroId   string `json:"outer_metro_id,omitempty" xml:"outer_metro_id,omitempty"`
    // 高德中心纬度坐标
    GaodeLatitude   string `json:"gaode_latitude,omitempty" xml:"gaode_latitude,omitempty"`
    // 高德中心经度坐标
    GaodeLongitude   string `json:"gaode_longitude,omitempty" xml:"gaode_longitude,omitempty"`
    // 是否删除 0 否 1是（默认0）
    IsDeleted   string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
    // 站点出口经纬度
    ExitPoi   string `json:"exit_poi,omitempty" xml:"exit_poi,omitempty"`
}
