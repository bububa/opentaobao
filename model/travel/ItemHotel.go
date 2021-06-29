package travel

// ItemHotel 
type ItemHotel struct {
    // 关联的套餐id
    RelatedPackageId   int64 `json:"related_package_id,omitempty" xml:"related_package_id,omitempty"`
    // 晚数
    HotelDays   int64 `json:"hotel_days,omitempty" xml:"hotel_days,omitempty"`
    // 酒店描述
    HotelDesc   string `json:"hotel_desc,omitempty" xml:"hotel_desc,omitempty"`
    // POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE
    PoiResource   string `json:"poi_resource,omitempty" xml:"poi_resource,omitempty"`
    // 酒店的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222
    Poi   string `json:"poi,omitempty" xml:"poi,omitempty"`
    // 必填，酒店房型
    HouseType   string `json:"house_type,omitempty" xml:"house_type,omitempty"`
    // 必填，酒店等级
    HotelLevel   string `json:"hotel_level,omitempty" xml:"hotel_level,omitempty"`
    // 必填，酒店名称
    CnName   string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
    // 必填，所在城市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
}
