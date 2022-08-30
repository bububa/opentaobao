package product

// GoodsPropertyValueDto 结构体
type GoodsPropertyValueDto struct {
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 属性值id，可不传
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 属性id
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}
