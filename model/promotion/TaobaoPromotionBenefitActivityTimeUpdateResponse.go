package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新关联活动有效时间 APIResponse
taobao.promotion.benefit.activity.time.update

更新关联权益的活动有效时间
*/
type TaobaoPromotionBenefitActivityTimeUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionBenefitActivityTimeUpdateResponse `json:"taobao_promotion_benefit_activity_time_update_response,omitempty"`
}

type TaobaoPromotionBenefitActivityTimeUpdateResponse struct {

    // 修改是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
