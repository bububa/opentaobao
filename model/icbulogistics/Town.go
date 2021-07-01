package icbulogistics

// Town 结构体
type Town struct {
	// 地址代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 地址代码
	AreaId string `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 地址名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
