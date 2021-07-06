package category

// TopPVPairDo 结构体
type TopPVPairDo struct {
	// 属性项名称
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
	// 属性值名称
	ValueName string `json:"value_name,omitempty" xml:"value_name,omitempty"`
	// 属性值ID
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 属性项ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}
