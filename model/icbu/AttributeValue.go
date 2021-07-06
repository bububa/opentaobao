package icbu

// AttributeValue 结构体
type AttributeValue struct {
	// 该属性值的子属性id
	ChildAttrs []int64 `json:"child_attrs,omitempty" xml:"child_attrs>int64,omitempty"`
	// 英文名字
	EnName string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// 属性id
	AttrId int64 `json:"attr_id,omitempty" xml:"attr_id,omitempty"`
	// 属性值id
	AttrValueId int64 `json:"attr_value_id,omitempty" xml:"attr_value_id,omitempty"`
	// 所属类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 是否SKU属性值
	SkuValue bool `json:"sku_value,omitempty" xml:"sku_value,omitempty"`
}
