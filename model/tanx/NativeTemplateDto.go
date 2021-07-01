package tanx

// NativeTemplateDto 结构体
type NativeTemplateDto struct {
	// 样式预览图片url
	Preview string `json:"preview,omitempty" xml:"preview,omitempty"`
	// 说明
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 模板ID
	TmplId int64 `json:"tmpl_id,omitempty" xml:"tmpl_id,omitempty"`
	// 区域列表
	Areas []NativeTemplateAreaDto `json:"areas,omitempty" xml:"areas>native_template_area_dto,omitempty"`
	// 模板支持的广告位尺寸
	Size string `json:"size,omitempty" xml:"size,omitempty"`
}
