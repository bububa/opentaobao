package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘专用单用户发放接口 APIResponse
taobao.mobile.promotion.benefit.activity.send

卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
*/
type TaobaoMobilePromotionBenefitActivitySendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"mobile_promotion_benefit_activity_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 权益发放结果
    
    SendResult   *SingleBenefitSendResult `json:"send_result,omitempty" xml:"