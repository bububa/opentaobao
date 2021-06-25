package icbu

// AttributeValueRequest 
type AttributeValueRequest struct {

    // 选填；需要过滤的属性值id
    AttributeValueId   []Number `json:"attribute_value_id,omitempty"`

    // 必填；要查询的属性值所属发布类目
    CatId   int64 `json:"cat_id,omitempty"`

    // 选填；需要过滤的属性
    AttributeId   []Number `json:"attribute_id,omitempty"`

}
