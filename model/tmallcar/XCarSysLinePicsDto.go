package tmallcar

// XCarSysLinePicsDto 结构体
type XCarSysLinePicsDto struct {
	// 具体颜色名称，除整体外观外，其他不区分颜色，此值为空，比如 黑色
	ColorName string `json:"color_name,omitempty" xml:"color_name,omitempty"`
	// 具体颜色色值，除整体外观外，其他不区分颜色，所以没有色值 比如 FFFF000FFF
	ColorValue string `json:"color_value,omitempty" xml:"color_value,omitempty"`
	// 图片列表List<URL>，最多100张,json格式 比如： ["www.baoud.com","www.alibb.com","www.accww.com"]
	Pics string `json:"pics,omitempty" xml:"pics,omitempty"`
	// 品牌属性id
	BrandPid int64 `json:"brand_pid,omitempty" xml:"brand_pid,omitempty"`
	// 品牌值id
	BrandVid int64 `json:"brand_vid,omitempty" xml:"brand_vid,omitempty"`
	// 车系属性id
	LinePid int64 `json:"line_pid,omitempty" xml:"line_pid,omitempty"`
	// 车系值id
	LineVid int64 `json:"line_vid,omitempty" xml:"line_vid,omitempty"`
	// 状态0.无效 1有效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// type(整体外观1、细节外观2、控件座椅3、中控区4、其他5)
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
