package alicom

// CatInfoVo 结构体
type CatInfoVo struct {
	// 手机归属区域
	ShowCatName string `json:"show_cat_name,omitempty" xml:"show_cat_name,omitempty"`
	// 手机归属运营商
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 手机归属城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 运营商
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
}
