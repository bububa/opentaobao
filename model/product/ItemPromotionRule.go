package product

// ItemPromotionRule 结构体
type ItemPromotionRule struct {
	// 规则名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 规则描叙信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 规则生效开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 规则生效结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 规则类型，常见有SKU锁定规则,下架锁定规则,库存减少锁定规则,库存禁止修改规则,一口价禁止修改规则
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}
