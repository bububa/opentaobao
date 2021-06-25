package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
手淘专用单用户发放接口 APIResponse
taobao.mobile.promotion.benefit.activity.send.share

卖家活动中需要通过该API来发放对应的权益。手淘专用、验证分享链路。
*/
type TaobaoMobilePromotionBenefitActivitySendShareAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMobilePromotionBenefitActivitySendShareResponse `json:"taobao_mobile_promotion_benefit_activity_send_share_response,omitempty"`
}

type TaobaoMobilePromotionBenefitActivitySendShareResponse struct {

    // 权益发放结果
    SendResult   *ShareBenefitSendResult `json:"send_result,omitempty"`

}
