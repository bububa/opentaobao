package promotion

// CouponTemplateParticipateConfig 
/* model for simplify = false
type CouponTemplateParticipateConfig struct {

    // 参与者列表
    
    ParticipateList  struct {
        LogicGroup  []LogicGroup `json:"logic_group,omitempty"`
    } `json:"participate_list,omitempty"`
    

}
*/

// CouponTemplateParticipateConfig 
type CouponTemplateParticipateConfig struct {

    // 参与者列表
    ParticipateList   []LogicGroup `json:"participate_list,omitempty"`

}
