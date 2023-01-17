package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityDeleteAPIRequest 删除关联的活动权益 API请求
// taobao.promotion.benefit.activity.delete
//
// 删除关联的活动权益
type TaobaoPromotionBenefitActivityDeleteAPIRequest struct {
	model.Params
	// ISV活动关联权益后获得的关联ID
	_relationId int64
}

// NewTaobaoPromotionBenefitActivityDeleteRequest 初始化TaobaoPromotionBenefitActivityDeleteAPIRequest对象
func NewTaobaoPromotionBenefitActivityDeleteRequest() *TaobaoPromotionBenefitActivityDeleteAPIRequest {
	return &TaobaoPromotionBenefitActivityDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionBenefitActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationId is RelationId Setter
// ISV活动关联权益后获得的关联ID
func (r *TaobaoPromotionBenefitActivityDeleteAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoPromotionBenefitActivityDeleteAPIRequest) GetRelationId() int64 {
	return r._relationId
}
