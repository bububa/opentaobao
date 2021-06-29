package wdk

// IndustryPropDO 
type IndustryPropDO struct {
    // 行业类型
    IndustryType   string `json:"industry_type,omitempty" xml:"industry_type,omitempty"`
    // 行业对应的属性
    Props   []PropDO `json:"props,omitempty" xml:"props>prop_do,omitempty"`
}
