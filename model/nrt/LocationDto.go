package nrt

// LocationDto 结构体
type LocationDto struct {
	// 城市名
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省份名
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 省份ID
	ProvId int64 `json:"prov_id,omitempty" xml:"prov_id,omitempty"`
	// 类型，1：国内，2：国外
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
