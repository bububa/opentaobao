package promotion

// BenefitDTO 
type BenefitDTO struct {
    // 权益code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 发放结束时间
    EndDate   string `json:"end_date,omitempty" xml:"end_date,omitempty"`
    // 权益描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 发放量
    Bestow   int64 `json:"bestow,omitempty" xml:"bestow,omitempty"`
    // 权益类型
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 发放开始时间
    StartDate   string `json:"start_date,omitempty" xml:"start_date,omitempty"`
    // 总库存
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 权益标签
    LabelCodes   string `json:"label_codes,omitempty" xml:"label_codes,omitempty"`
    // 扩展信息
    Feature   *Feature `json:"feature,omitempty" xml:"feature,omitempty"`
    // 权益名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 权益状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
}
