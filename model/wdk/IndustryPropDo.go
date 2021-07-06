package wdk

// IndustryPropDo 结构体
type IndustryPropDo struct {
	// 行业对应的属性
	Props []PropDo `json:"props,omitempty" xml:"props>prop_do,omitempty"`
	// 行业类型
	IndustryType string `json:"industry_type,omitempty" xml:"industry_type,omitempty"`
}
