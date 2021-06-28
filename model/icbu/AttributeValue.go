package icbu

// AttributeValue 
/* model for simplify = false
type AttributeValue struct {

    // 属性id
    
    AttrId   int64 `json:"attr_id,omitempty"`
    

    // 英文名字
    
    EnName   string `json:"en_name,omitempty"`
    

    // 属性值id
    
    AttrValueId   int64 `json:"attr_value_id,omitempty"`
    

    // 该属性值的子属性id
    
    ChildAttrs  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"child_attrs,omitempty"`
    

    // 所属类目id
    
    CatId   int64 `json:"cat_id,omitempty"`
    

    // 是否SKU属性值
    
    SkuValue   bool `json:"sku_value,omitempty"`
    

}
*/

// AttributeValue 
type AttributeValue struct {

    // 属性id
    AttrId   int64 `json:"attr_id,omitempty"`

    // 英文名字
    EnName   string `json:"en_name,omitempty"`

    // 属性值id
    AttrValueId   int64 `json:"attr_value_id,omitempty"`

    // 该属性值的子属性id
    ChildAttrs   []int64 `json:"child_attrs,omitempty"`

    // 所属类目id
    CatId   int64 `json:"cat_id,omitempty"`

    // 是否SKU属性值
    SkuValue   bool `json:"sku_value,omitempty"`

}
