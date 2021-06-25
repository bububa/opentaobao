package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关联活动权益 APIResponse
taobao.promotion.benefit.activity.relation

卖家活动中需要通过该API来关联的对应的权益。
*/
type TaobaoPromotionBenefitActivityRelationAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionBenefitActivityRelationResponse `json:"taobao_promotion_benefit_activity_relation_response,omitempty"`
}

type TaobaoPromotionBenefitActivityRelationResponse struct {

    // 活动关联ID
    RelationId   int64 `json:"relation_id,omitempty"`

    // 请求是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
