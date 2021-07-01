package store

// PoiInfoDto 
type PoiInfoDto struct {
    // 地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 地址别名
    AddressAlias   string `json:"address_alias,omitempty" xml:"address_alias,omitempty"`
    // 英文地址
    AddressEn   string `json:"address_en,omitempty" xml:"address_en,omitempty"`
    // 当地语言地址
    AddressLocal   string `json:"address_local,omitempty" xml:"address_local,omitempty"`
    // 铺位号
    Bunk   string `json:"bunk,omitempty" xml:"bunk,omitempty"`
    // 营业面积
    BusinessArea   string `json:"business_area,omitempty" xml:"business_area,omitempty"`
    // 市
    City   int64 `json:"city,omitempty" xml:"city,omitempty"`
    // 省
    CityInt   int64 `json:"city_int,omitempty" xml:"city_int,omitempty"`
    // 市名
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    // 洲
    Continent   int64 `json:"continent,omitempty" xml:"continent,omitempty"`
    // 洲名
    ContinentName   string `json:"continent_name,omitempty" xml:"continent_name,omitempty"`
    // 国家
    Country   int64 `json:"country,omitempty" xml:"country,omitempty"`
    // 国家名
    CountryName   string `json:"country_name,omitempty" xml:"country_name,omitempty"`
    // 区域
    District   int64 `json:"district,omitempty" xml:"district,omitempty"`
    // 区域
    DistrictInt   int64 `json:"district_int,omitempty" xml:"district_int,omitempty"`
    // 区域名
    DistrictName   string `json:"district_name,omitempty" xml:"district_name,omitempty"`
    // 行政区code类型
    DivisionType   int64 `json:"division_type,omitempty" xml:"division_type,omitempty"`
    // 楼层
    Floor   string `json:"floor,omitempty" xml:"floor,omitempty"`
    // 楼层名
    FullAddress   string `json:"full_address,omitempty" xml:"full_address,omitempty"`
    // 邮政编码
    PostCode   string `json:"post_code,omitempty" xml:"post_code,omitempty"`
    // 经度
    Posx   string `json:"posx,omitempty" xml:"posx,omitempty"`
    // 纬度
    Posy   string `json:"posy,omitempty" xml:"posy,omitempty"`
    // 省
    Prov   int64 `json:"prov,omitempty" xml:"prov,omitempty"`
    // 省
    ProvInt   int64 `json:"prov_int,omitempty" xml:"prov_int,omitempty"`
    // 省名
    ProvName   string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
    // 街道
    Town   int64 `json:"town,omitempty" xml:"town,omitempty"`
    // 街道
    TownInt   int64 `json:"town_int,omitempty" xml:"town_int,omitempty"`
    // 街道名
    TownName   string `json:"town_name,omitempty" xml:"town_name,omitempty"`
}
