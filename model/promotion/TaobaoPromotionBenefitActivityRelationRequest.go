package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关联活动权益 APIRequest
taobao.promotion.benefit.activity.relation

卖家活动中需要通过该API来关联的对应的权益。
*/
type TaobaoPromotionBenefitActivityRelationRequest struct {
    model.Params

    // 活动关联权益请求参数
    relationRequest   *RelationActivityBenefitRequest 

}

func NewTaobaoPromotionBenefitActivityRelationRequest() *TaobaoPromotionBenefitActivityRelationRequest{
    return &TaobaoPromotionBenefitActivityRelationRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionBenefitActivityRelationRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.relation"
}

func (r TaobaoPromotionBenefitActivityRelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionBenefitActivityRelationRequest) SetRelationRequest(relationRequest *RelationActivityBenefitRequest) error {
    r.relationRequest = relationRequest
    r.Set("relation_request", relationRequest)
    return nil
}

func (r TaobaoPromotionBenefitActivityRelationRequest) GetRelationRequest() *RelationActivityBenefitRequest {
    return r.relationRequest
}

