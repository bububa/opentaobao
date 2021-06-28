package wdk

// PropDo 
type PropDo struct {

    // 行业属性
    
    Key   string `json:"key,omitempty" xml:"key,omitempty"`
    

    // 类目属性值
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    

    // 是否移除该属性
    
    RemoveOpt   bool `json:"remove_opt,omitempty" xml:"remove_opt,omitempty"`
    

}
