package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关联活动权益 API返回值 
taobao.promotion.benefit.activity.relation

卖家活动中需要通过该API来关联的对应的权益。
*/
type TaobaoPromotionBenefitActivityRelationAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionBenefitActivityRelationResponse
}

// 关联活动权益 成功返回结果
type TaobaoPromotionBenefitActivityRelationResponse struct {
    XMLName xml.Name `xml:"promotion_benefit_activity_relation_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 活动关联ID
    RelationId   int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
    // 请求是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
