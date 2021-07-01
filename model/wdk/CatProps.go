package wdk

// CatProps 结构体
type CatProps struct {
	// 类目属性ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 类目属性名称
	PropertyText string `json:"property_text,omitempty" xml:"property_text,omitempty"`
	// 类目值ID
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 类目值名称
	ValueText string `json:"value_text,omitempty" xml:"value_text,omitempty"`
}
