package idleisv

// YoupinCpvQry 结构体
type YoupinCpvQry struct {
	// 属性值id
	ValueId string `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 属性id
	PropertyId string `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 类目id
	ChannelCatId string `json:"channel_cat_id,omitempty" xml:"channel_cat_id,omitempty"`
}
