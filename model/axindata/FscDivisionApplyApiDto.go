package axindata

// FscDivisionApplyApiDto 结构体
type FscDivisionApplyApiDto struct {
	// 行政区划外部编号（供应商侧编号）
	DivisionOuterId string `json:"division_outer_id,omitempty" xml:"division_outer_id,omitempty"`
	// 行政区划名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 行政区划英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 行政区划父级ID
	ParentId string `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 是否境外
	Abroad bool `json:"abroad,omitempty" xml:"abroad,omitempty"`
}
