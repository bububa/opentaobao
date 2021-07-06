package icbu

// AttributeValueRequest 结构体
type AttributeValueRequest struct {
	// 选填；需要过滤的属性值id
	AttributeValueId []int64 `json:"attribute_value_id,omitempty" xml:"attribute_value_id>int64,omitempty"`
	// 选填；需要过滤的属性
	AttributeId []int64 `json:"attribute_id,omitempty" xml:"attribute_id>int64,omitempty"`
	// 必填；要查询的属性值所属发布类目
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}
