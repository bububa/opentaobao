package promotion

// CouponTemplateParticipateConfig 
type CouponTemplateParticipateConfig struct {
    // 参与者列表
    ParticipateList   []LogicGroup `json:"participate_list,omitempty" xml:"participate_list>logic_group,omitempty"`
}
