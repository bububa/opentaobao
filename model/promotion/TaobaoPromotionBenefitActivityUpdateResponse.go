package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改关联的活动权益 APIResponse
taobao.promotion.benefit.activity.update

修改卖家活动中关联的对应的权益。
*/
type TaobaoPromotionBenefitActivityUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionBenefitActivityUpdateResponse `json:"promotion_benefit_activity_update_response,omitempty"` 
    TaobaoPromotionBenefitActivityUpdateResponse
}

/* model for simplify = false
type TaobaoPromotionBenefitActivityUpdateResponse struct {

    // 更新是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoPromotionBenefitActivityUpdateResponse struct {

    // 更新是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
