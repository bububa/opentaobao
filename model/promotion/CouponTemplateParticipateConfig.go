package promotion

// CouponTemplateParticipateConfig 结构体
type CouponTemplateParticipateConfig struct {
	// 参与者列表
	ParticipateList []LogicGroup `json:"participate_list,omitempty" xml:"participate_list>logic_group,omitempty"`
}
