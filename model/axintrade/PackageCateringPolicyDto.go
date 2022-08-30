package axintrade

// PackageCateringPolicyDto 结构体
type PackageCateringPolicyDto struct {
	// 餐饮使用规则名称
	CateringRuleName string `json:"catering_rule_name,omitempty" xml:"catering_rule_name,omitempty"`
	// 就餐开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 就餐截止时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 就餐方式
	CateringWay string `json:"catering_way,omitempty" xml:"catering_way,omitempty"`
	// 餐饮规则非结构化描述
	CateringDescription string `json:"catering_description,omitempty" xml:"catering_description,omitempty"`
	// 其他产品使用规则名称
	OtherRuleName string `json:"other_rule_name,omitempty" xml:"other_rule_name,omitempty"`
	// 其他产品非结构化描述
	OtherDescription string `json:"other_description,omitempty" xml:"other_description,omitempty"`
}
