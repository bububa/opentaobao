package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
手淘专用单用户发放接口 APIResponse
taobao.mobile.promotion.benefit.activity.send

卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
*/
type TaobaoMobilePromotionBenefitActivitySendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMobilePromotionBenefitActivitySendResponse `json:"mobile_promotion_benefit_activity_send_response,omitempty"` 
    TaobaoMobilePromotionBenefitActivitySendResponse
}

/* model for simplify = false
type TaobaoMobilePromotionBenefitActivitySendResponse struct {

    // 权益发放结果
    
    SendResult  *struct {
        SingleBenefitSendResult  *SingleBenefitSendResult `json:"single_benefit_send_result,omitempty"`
    } `json:"send_result,omitempty"`
    

}
*/

type TaobaoMobilePromotionBenefitActivitySendResponse struct {

    // 权益发放结果
    SendResult   *SingleBenefitSendResult `json:"send_result,omitempty"`

}
