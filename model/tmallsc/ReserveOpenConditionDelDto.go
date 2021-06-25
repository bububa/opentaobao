package tmallsc

// ReserveOpenConditionDelDto 
type ReserveOpenConditionDelDto struct {

    // 城市id
    CityId   int64 `json:"city_id,omitempty"`

    // 类目id
    CategoryId   int64 `json:"category_id,omitempty"`

    // 品牌id
    BrandId   int64 `json:"brand_id,omitempty"`

    // 服务code
    ServiceCode   string `json:"service_code,omitempty"`

    // 省份id
    ProvinceId   int64 `json:"province_id,omitempty"`

    // 区域集合ids
    AreaIds   string `json:"area_ids,omitempty"`

}
