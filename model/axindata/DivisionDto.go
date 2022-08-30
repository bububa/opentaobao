package axindata

// DivisionDto 结构体
type DivisionDto struct {
	// 1-国家 2-省 3-城市
	Level string `json:"level,omitempty" xml:"level,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 英文名称
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 代码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
