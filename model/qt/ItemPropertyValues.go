package qt

// ItemPropertyValues 
/* model for simplify = false
type ItemPropertyValues struct {

    // 服务属性id
    
    PropertyId   int64 `json:"property_id,omitempty"`
    

    // 属性名称
    
    PropertyName   string `json:"property_name,omitempty"`
    

    // 属性值列表.
    
    PropertyValues  struct {
        String  []string `json:"string,omitempty"`
    } `json:"property_values,omitempty"`
    

}
*/

// ItemPropertyValues 
type ItemPropertyValues struct {

    // 服务属性id
    PropertyId   int64 `json:"property_id,omitempty"`

    // 属性名称
    PropertyName   string `json:"property_name,omitempty"`

    // 属性值列表.
    PropertyValues   []string `json:"property_values,omitempty"`

}
