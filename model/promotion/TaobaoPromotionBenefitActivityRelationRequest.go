package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关联活动权益 API请求
taobao.promotion.benefit.activity.relation

卖家活动中需要通过该API来关联的对应的权益。
*/
type TaobaoPromotionBenefitActivityRelationRequest struct {
    model.Params
    // 活动关联权益请求参数
    relationRequest   *RelationActivityBenefitRequest
}

// 初始化TaobaoPromotionBenefitActivityRelationRequest对象
func NewTaobaoPromotionBenefitActivityRelationRequest() *TaobaoPromotionBenefitActivityRelationRequest{
    return &TaobaoPromotionBenefitActivityRelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityRelationRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.relation"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityRelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RelationRequest Setter
// 活动关联权益请求参数
func (r *TaobaoPromotionBenefitActivityRelationRequest) SetRelationRequest(relationRequest *RelationActivityBenefitRequest) error {
    r.relationRequest = relationRequest
    r.Set("relation_request", relationRequest)
    return nil
}

// RelationRequest Getter
func (r TaobaoPromotionBenefitActivityRelationRequest) GetRelationRequest() *RelationActivityBenefitRequest {
    return r.relationRequest
}
