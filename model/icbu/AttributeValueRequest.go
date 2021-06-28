package icbu

// AttributeValueRequest 
/* model for simplify = false
type AttributeValueRequest struct {

    // 选填；需要过滤的属性值id
    
    AttributeValueId  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"attribute_value_id,omitempty"`
    

    // 必填；要查询的属性值所属发布类目
    
    CatId   int64 `json:"cat_id,omitempty"`
    

    // 选填；需要过滤的属性
    
    AttributeId  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"attribute_id,omitempty"`
    

}
*/

// AttributeValueRequest 
type AttributeValueRequest struct {

    // 选填；需要过滤的属性值id
    AttributeValueId   []int64 `json:"attribute_value_id,omitempty"`

    // 必填；要查询的属性值所属发布类目
    CatId   int64 `json:"cat_id,omitempty"`

    // 选填；需要过滤的属性
    AttributeId   []int64 `json:"attribute_id,omitempty"`

}
