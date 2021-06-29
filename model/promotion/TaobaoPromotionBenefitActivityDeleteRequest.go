package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关联的活动权益 API请求
taobao.promotion.benefit.activity.delete

删除关联的活动权益
*/
type TaobaoPromotionBenefitActivityDeleteRequest struct {
    model.Params
    // ISV活动关联权益后获得的关联ID
    relationId   int64
}

// 初始化TaobaoPromotionBenefitActivityDeleteRequest对象
func NewTaobaoPromotionBenefitActivityDeleteRequest() *TaobaoPromotionBenefitActivityDeleteRequest{
    return &TaobaoPromotionBenefitActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RelationId Setter
// ISV活动关联权益后获得的关联ID
func (r *TaobaoPromotionBenefitActivityDeleteRequest) SetRelationId(relationId int64) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

// RelationId Getter
func (r TaobaoPromotionBenefitActivityDeleteRequest) GetRelationId() int64 {
    return r.relationId
}
