package promotion

// OrightDto 结构体
type OrightDto struct {
	// 奖品id
	PrizeId int64 `json:"prize_id,omitempty" xml:"prize_id,omitempty"`
	// 模板名称
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 权益内容
	BenefitName string `json:"benefit_name,omitempty" xml:"benefit_name,omitempty"`
	// 权益类型名称
	RightTypeName string `json:"right_type_name,omitempty" xml:"right_type_name,omitempty"`
	// 权益类型id
	RightTypeId int64 `json:"right_type_id,omitempty" xml:"right_type_id,omitempty"`
	// 开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 总数
	PrizeQuantity int64 `json:"prize_quantity,omitempty" xml:"prize_quantity,omitempty"`
	// 可发放数
	RemainPrizeQuantity int64 `json:"remain_prize_quantity,omitempty" xml:"remain_prize_quantity,omitempty"`
	// 概率
	Probability string `json:"probability,omitempty" xml:"probability,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 使用开始时间
	UseStartTime string `json:"use_start_time,omitempty" xml:"use_start_time,omitempty"`
	// 使用结束时间
	UseEndTime string `json:"use_end_time,omitempty" xml:"use_end_time,omitempty"`
	// 门槛
	Condition string `json:"condition,omitempty" xml:"condition,omitempty"`
	// 扩展参数
	ExtAttribute string `json:"ext_attribute,omitempty" xml:"ext_attribute,omitempty"`
}
