package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionbenefitactivityrelationAPIRequest 关联活动权益 API请求
// taobao.promotion.benefit.activity.relation
//
// 卖家活动中需要通过该API来关联的对应的权益。
type TaobaopromotionbenefitactivityrelationAPIRequest struct {
	model.Params
	// 活动关联权益请求参数
	_relationRequest *RelationActivityBenefitRequest
}

// NewTaobaopromotionbenefitactivityrelationRequest 初始化TaobaopromotionbenefitactivityrelationAPIRequest对象
func NewTaobaopromotionbenefitactivityrelationRequest() *TaobaopromotionbenefitactivityrelationAPIRequest {
	return &TaobaopromotionbenefitactivityrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionbenefitactivityrelationAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionbenefitactivityrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionbenefitactivityrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationRequest is RelationRequest Setter
// 活动关联权益请求参数
func (r *TaobaopromotionbenefitactivityrelationAPIRequest) SetRelationRequest(_relationRequest *RelationActivityBenefitRequest) error {
	r._relationRequest = _relationRequest
	r.Set("relation_request", _relationRequest)
	return nil
}

// GetRelationRequest RelationRequest Getter
func (r TaobaopromotionbenefitactivityrelationAPIRequest) GetRelationRequest() *RelationActivityBenefitRequest {
	return r._relationRequest
}
