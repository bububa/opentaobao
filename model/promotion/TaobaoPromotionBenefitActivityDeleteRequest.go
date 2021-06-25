package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关联的活动权益 APIRequest
taobao.promotion.benefit.activity.delete

删除关联的活动权益
*/
type TaobaoPromotionBenefitActivityDeleteRequest struct {
    model.Params

    // ISV活动关联权益后获得的关联ID
    relationId   int64 

}

func NewTaobaoPromotionBenefitActivityDeleteRequest() *TaobaoPromotionBenefitActivityDeleteRequest{
    return &TaobaoPromotionBenefitActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionBenefitActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.delete"
}

func (r TaobaoPromotionBenefitActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionBenefitActivityDeleteRequest) SetRelationId(relationId int64) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

func (r TaobaoPromotionBenefitActivityDeleteRequest) GetRelationId() int64 {
    return r.relationId
}

