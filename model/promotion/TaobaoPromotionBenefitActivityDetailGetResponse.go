package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
活动关联的权益详情获取 APIResponse
taobao.promotion.benefit.activity.detail.get

活动关联的权益详情获取
*/
type TaobaoPromotionBenefitActivityDetailGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionBenefitActivityDetailGetResponse `json:"taobao_promotion_benefit_activity_detail_get_response,omitempty"`
}

type TaobaoPromotionBenefitActivityDetailGetResponse struct {

    // 查询是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 活动关联的权益详情列表
    RelationBenefitDetails   string `json:"relation_benefit_details,omitempty"`

}
