package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关联的活动权益 API返回值 
taobao.promotion.benefit.activity.delete

删除关联的活动权益
*/
type TaobaoPromotionBenefitActivityDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionBenefitActivityDeleteResponse
}

// 删除关联的活动权益 成功返回结果
type TaobaoPromotionBenefitActivityDeleteResponse struct {
    XMLName xml.Name `xml:"promotion_benefit_activity_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 删除是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
