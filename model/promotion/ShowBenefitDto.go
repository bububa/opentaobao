package promotion

// ShowBenefitDto 结构体
type ShowBenefitDto struct {
	// 权益code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 权益类型
	TypeDesc string `json:"type_desc,omitempty" xml:"type_desc,omitempty"`
	// 权益展示面额单位
	DisplayAmountUnit string `json:"display_amount_unit,omitempty" xml:"display_amount_unit,omitempty"`
	// 权益发放结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 权益标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 权益类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 权益扩展信息
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 权益是否可领
	CanWin bool `json:"can_win,omitempty" xml:"can_win,omitempty"`
	// 相对使用时间类型
	IntervalTimeUnit string `json:"interval_time_unit,omitempty" xml:"interval_time_unit,omitempty"`
	// 展示门槛
	DisplayStartFee string `json:"display_start_fee,omitempty" xml:"display_start_fee,omitempty"`
	// 权益发放类型
	SendMode string `json:"send_mode,omitempty" xml:"send_mode,omitempty"`
	// 权益生命周期
	SendLifeCycleState string `json:"send_life_cycle_state,omitempty" xml:"send_life_cycle_state,omitempty"`
	// 权益面额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 是否测试权益
	Test bool `json:"test,omitempty" xml:"test,omitempty"`
	// 是否已领
	HadWin bool `json:"had_win,omitempty" xml:"had_win,omitempty"`
	// 权益规则列表
	ShowRules []ShowRuleDto `json:"show_rules,omitempty" xml:"show_rules>show_rule_dto,omitempty"`
	// 面额单位
	AmountUnit string `json:"amount_unit,omitempty" xml:"amount_unit,omitempty"`
	// 是否有库存
	HasInventory bool `json:"has_inventory,omitempty" xml:"has_inventory,omitempty"`
	// 展示面额
	DisplayAmount string `json:"display_amount,omitempty" xml:"display_amount,omitempty"`
	// 使用时间类型
	EffectiveTimeMode string `json:"effective_time_mode,omitempty" xml:"effective_time_mode,omitempty"`
	// 素材
	Material string `json:"material,omitempty" xml:"material,omitempty"`
	// 相对使用时间长度
	EffectiveInterval int64 `json:"effective_interval,omitempty" xml:"effective_interval,omitempty"`
	// 加密动态面额参数
	EncryptedDynamicInfo string `json:"encrypted_dynamic_info,omitempty" xml:"encrypted_dynamic_info,omitempty"`
	// 门槛
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 权益安全码
	Asac string `json:"asac,omitempty" xml:"asac,omitempty"`
	// 发放开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 绝对使用时间开始
	EffectiveStart string `json:"effective_start,omitempty" xml:"effective_start,omitempty"`
	// 绝对使用时间结束
	EffectiveEnd string `json:"effective_end,omitempty" xml:"effective_end,omitempty"`
}
