package tbitem

// Location 结构体
type Location struct {
	// 所在城市（中文名称）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 所在省份（中文名称）
	State string `json:"state,omitempty" xml:"state,omitempty"`
}
