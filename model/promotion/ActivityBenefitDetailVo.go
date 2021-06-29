package promotion

// ActivityBenefitDetailVo 
type ActivityBenefitDetailVo struct {
    // 权益类型
    BenefitType   string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
    // 权益ID
    BenefitId   int64 `json:"benefit_id,omitempty" xml:"benefit_id,omitempty"`
    // 权益标识
    ConfigId   int64 `json:"config_id,omitempty" xml:"config_id,omitempty"`
}
