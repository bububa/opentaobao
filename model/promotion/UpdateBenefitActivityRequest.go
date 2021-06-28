package promotion

// UpdateBenefitActivityRequest 
/* model for simplify = false
type UpdateBenefitActivityRequest struct {

    // 同步权益活动的概述信息，方便卖家后台查看
    
    BenefitActivityVo  *struct {
        UpdateBenefitActivityVo  *UpdateBenefitActivityVo `json:"update_benefit_activity_vo,omitempty"`
    } `json:"benefit_activity_vo,omitempty"`
    

    // 需要删除的已经关联的权益
    
    DeleteDetailVos  struct {
        DeleteActivityBenefitDetailVo  []DeleteActivityBenefitDetailVo `json:"delete_activity_benefit_detail_vo,omitempty"`
    } `json:"delete_detail_vos,omitempty"`
    

    // 活动关联的权益信息列表，可以从权益选择器API中获取
    
    AddDetailVos  struct {
        ActivityBenefitDetailVo  []ActivityBenefitDetailVo `json:"activity_benefit_detail_vo,omitempty"`
    } `json:"add_detail_vos,omitempty"`
    

}
*/

// UpdateBenefitActivityRequest 
type UpdateBenefitActivityRequest struct {

    // 同步权益活动的概述信息，方便卖家后台查看
    BenefitActivityVo   *UpdateBenefitActivityVo `json:"benefit_activity_vo,omitempty"`

    // 需要删除的已经关联的权益
    DeleteDetailVos   []DeleteActivityBenefitDetailVo `json:"delete_detail_vos,omitempty"`

    // 活动关联的权益信息列表，可以从权益选择器API中获取
    AddDetailVos   []ActivityBenefitDetailVo `json:"add_detail_vos,omitempty"`

}
