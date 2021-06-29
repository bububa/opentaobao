package maitix

// MoaAddressInfo 
type MoaAddressInfo struct {
    // 国家ID，目前只支持中国，传1-如果是快递票必填
    CountryId   int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
    // 省ID，国标-如果是快递票必填
    ProvinceId   int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
    // 市ID，国标，快递票该字段必须填
    CityId   int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
    // 区域ID，国标-如果是快递票必填
    AreaId   int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}
