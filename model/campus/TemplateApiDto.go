package campus

// TemplateApiDto 结构体
type TemplateApiDto struct {
	// 参数点集合
	PropertyList []PropertyApiDto `json:"property_list,omitempty" xml:"property_list>property_api_dto,omitempty"`
	// 模板编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 模板名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 模板id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
