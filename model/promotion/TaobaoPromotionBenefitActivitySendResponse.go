package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动权益发放接口 APIResponse
taobao.promotion.benefit.activity.send

活动权益发放接口，用于卖家针对活动进行权益发放
*/
type TaobaoPromotionBenefitActivitySendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_benefit_activity_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"