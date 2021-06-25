package promotion

// OrightDto 
type OrightDto struct {

    // 奖品id
    PrizeId   int64 `json:"prize_id,omitempty"`

    // 模板名称
    TemplateName   string `json:"template_name,omitempty"`

    // 权益内容
    BenefitName   string `json:"benefit_name,omitempty"`

    // 权益类型名称
    RightTypeName   string `json:"right_type_name,omitempty"`

    // 权益类型id
    RightTypeId   int64 `json:"right_type_id,omitempty"`

    // 开始时间
    StartDate   string `json:"start_date,omitempty"`

    // 结束时间
    EndDate   string `json:"end_date,omitempty"`

    // 总数
    PrizeQuantity   int64 `json:"prize_quantity,omitempty"`

    // 可发放数
    RemainPrizeQuantity   int64 `json:"remain_prize_quantity,omitempty"`

    // 概率
    Probability   string `json:"probability,omitempty"`

    // 金额
    Amount   string `json:"amount,omitempty"`

    // 使用开始时间
    UseStartTime   string `json:"use_start_time,omitempty"`

    // 使用结束时间
    UseEndTime   string `json:"use_end_time,omitempty"`

    // 门槛
    Condition   string `json:"condition,omitempty"`

    // 扩展参数
    ExtAttribute   string `json:"ext_attribute,omitempty"`

}
