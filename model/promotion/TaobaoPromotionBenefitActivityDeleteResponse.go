package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关联的活动权益 APIResponse
taobao.promotion.benefit.activity.delete

删除关联的活动权益
*/
type TaobaoPromotionBenefitActivityDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_benefit_activity_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"