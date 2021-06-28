package wdk

// PropDo 
/* model for simplify = false
type PropDo struct {

    // 行业属性
    
    Key   string `json:"key,omitempty"`
    

    // 类目属性值
    
    Value   string `json:"value,omitempty"`
    

    // 是否移除该属性
    
    RemoveOpt   bool `json:"remove_opt,omitempty"`
    

}
*/

// PropDo 
type PropDo struct {

    // 行业属性
    Key   string `json:"key,omitempty"`

    // 类目属性值
    Value   string `json:"value,omitempty"`

    // 是否移除该属性
    RemoveOpt   bool `json:"remove_opt,omitempty"`

}
