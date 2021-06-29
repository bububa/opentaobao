package idleisv

// IdleNewPubValueDO 
type IdleNewPubValueDO struct {
    // 属性id
    PropertyId   string `json:"property_id,omitempty" xml:"property_id,omitempty"`
    // 属性名称
    PropertyName   string `json:"property_name,omitempty" xml:"property_name,omitempty"`
    // 渠道类目id
    ChannelCatId   string `json:"channel_cat_id,omitempty" xml:"channel_cat_id,omitempty"`
    // 值id
    ValueId   string `json:"value_id,omitempty" xml:"value_id,omitempty"`
    // 值名称
    ValueName   string `json:"value_name,omitempty" xml:"value_name,omitempty"`
}
