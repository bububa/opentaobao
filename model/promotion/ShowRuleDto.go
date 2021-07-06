package promotion

// ShowRuleDto 结构体
type ShowRuleDto struct {
	// 规则扩展信息
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 规则类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 规则是否通过
	Passed bool `json:"passed,omitempty" xml:"passed,omitempty"`
}
