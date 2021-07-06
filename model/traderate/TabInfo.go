package traderate

// TabInfo 结构体
type TabInfo struct {
	// tab筛选信息Code，查询时使用
	TabCode string `json:"tab_code,omitempty" xml:"tab_code,omitempty"`
	// 包含的数量
	TabDetail string `json:"tab_detail,omitempty" xml:"tab_detail,omitempty"`
	// tab名称
	TabName string `json:"tab_name,omitempty" xml:"tab_name,omitempty"`
	// 属性（正面负面）
	Attitude int64 `json:"attitude,omitempty" xml:"attitude,omitempty"`
	// 会否选中
	IsClick bool `json:"is_click,omitempty" xml:"is_click,omitempty"`
}
