package axindata

// FscTripDivisionApiDto 结构体
type FscTripDivisionApiDto struct {
	// 行政区划名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 行政区划英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 目的地ID
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
	// 行政区划级别
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 行政区划父级ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 是否境外
	Abroad bool `json:"abroad,omitempty" xml:"abroad,omitempty"`
}
