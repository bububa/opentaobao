package tmallsc

// ReserveOpenConditionDTO 
type ReserveOpenConditionDTO struct {

    // 类目id
    
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

    // 品牌id
    
    BrandId   int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 服务code
    
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    

    // 排除的区域id
    
    ExcludeAreaIds   string `json:"exclude_area_ids,omitempty" xml:"exclude_area_ids,omitempty"`
    

    // 区域ids集合
    
    AreaIds   string `json:"area_ids,omitempty" xml:"area_ids,omitempty"`
    

    // 城市id
    
    CityId   int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
    

    // 身份id
    
    ProvinceId   int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
    

}
