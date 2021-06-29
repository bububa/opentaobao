package store

// TaobaoPlaceStorerelatesubGetT 
type TaobaoPlaceStorerelatesubGetT struct {
    // 门店Id
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 门店名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 门店分名称
    SubName   string `json:"sub_name,omitempty" xml:"sub_name,omitempty"`
    // 详细地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 省
    ProvName   string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
    // 市
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    // 区
    DistrictName   string `json:"district_name,omitempty" xml:"district_name,omitempty"`
    // 省
    ProvCode   int64 `json:"prov_code,omitempty" xml:"prov_code,omitempty"`
    // 市
    CityCode   int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
    // 区
    DistrictCode   int64 `json:"district_code,omitempty" xml:"district_code,omitempty"`
}
