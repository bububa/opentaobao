package tmallsc

// ReserveOpenConditionDTO 
/* model for simplify = false
type ReserveOpenConditionDTO struct {

    // 类目id
    
    CategoryId   int64 `json:"category_id,omitempty"`
    

    // 品牌id
    
    BrandId   int64 `json:"brand_id,omitempty"`
    

    // 服务code
    
    ServiceCode   string `json:"service_code,omitempty"`
    

    // 排除的区域id
    
    ExcludeAreaIds   string `json:"exclude_area_ids,omitempty"`
    

    // 区域ids集合
    
    AreaIds   string `json:"area_ids,omitempty"`
    

    // 城市id
    
    CityId   int64 `json:"city_id,omitempty"`
    

    // 身份id
    
    ProvinceId   int64 `json:"province_id,omitempty"`
    

}
*/

// ReserveOpenConditionDTO 
type ReserveOpenConditionDTO struct {

    // 类目id
    CategoryId   int64 `json:"category_id,omitempty"`

    // 品牌id
    BrandId   int64 `json:"brand_id,omitempty"`

    // 服务code
    ServiceCode   string `json:"service_code,omitempty"`

    // 排除的区域id
    ExcludeAreaIds   string `json:"exclude_area_ids,omitempty"`

    // 区域ids集合
    AreaIds   string `json:"area_ids,omitempty"`

    // 城市id
    CityId   int64 `json:"city_id,omitempty"`

    // 身份id
    ProvinceId   int64 `json:"province_id,omitempty"`

}
