package store

// UpdateStoreExtendDto 
type UpdateStoreExtendDto struct {
    // 需要添加的TAGS
    AddTags   []int64 `json:"add_tags,omitempty" xml:"add_tags>int64,omitempty"`
    // 需要修改的attribute对应的key
    AttributeKey   string `json:"attribute_key,omitempty" xml:"attribute_key,omitempty"`
    // 需要修改的attribute对应的value
    AttributeValue   string `json:"attribute_value,omitempty" xml:"attribute_value,omitempty"`
    // 需要修改的bizAttribute对应的key
    BizAttributeKey   string `json:"biz_attribute_key,omitempty" xml:"biz_attribute_key,omitempty"`
    // 需要修改的bizAttribute对应的value
    BizAttributeValue   string `json:"biz_attribute_value,omitempty" xml:"biz_attribute_value,omitempty"`
    // 需要删除的TAGS
    RemoveTags   []int64 `json:"remove_tags,omitempty" xml:"remove_tags>int64,omitempty"`
    // 门店id
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
