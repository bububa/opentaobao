package wdk

// IndustryPropDo 
/* model for simplify = false
type IndustryPropDo struct {

    // 行业类型
    
    IndustryType   string `json:"industry_type,omitempty"`
    

    // 行业对应的属性
    
    Props  struct {
        PropDo  []PropDo `json:"prop_do,omitempty"`
    } `json:"props,omitempty"`
    

}
*/

// IndustryPropDo 
type IndustryPropDo struct {

    // 行业类型
    IndustryType   string `json:"industry_type,omitempty"`

    // 行业对应的属性
    Props   []PropDo `json:"props,omitempty"`

}
