package promotion

// RelationActivityBenefitRequest 
/* model for simplify = false
type RelationActivityBenefitRequest struct {

    // 同步权益活动的概述信息，用于卖家后台查看
    
    BenefitActivityVo  *struct {
        BenefitActivityVo  *BenefitActivityVo `json:"benefit_activity_vo,omitempty"`
    } `json:"benefit_activity_vo,omitempty"`
    

    // 活动关联的权益信息，可以从权益选择器API中获取
    
    AddDetailVos  struct {
        ActivityBenefitDetailVo  []ActivityBenefitDetailVo `json:"activity_benefit_detail_vo,omitempty"`
    } `json:"add_detail_vos,omitempty"`
    

}
*/

// RelationActivityBenefitRequest 
type RelationActivityBenefitRequest struct {

    // 同步权益活动的概述信息，用于卖家后台查看
    BenefitActivityVo   *BenefitActivityVo `json:"benefit_activity_vo,omitempty"`

    // 活动关联的权益信息，可以从权益选择器API中获取
    AddDetailVos   []ActivityBenefitDetailVo `json:"add_detail_vos,omitempty"`

}
