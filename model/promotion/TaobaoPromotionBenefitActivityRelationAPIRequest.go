package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityRelationAPIRequest 关联活动权益 API请求
// taobao.promotion.benefit.activity.relation
//
// 卖家活动中需要通过该API来关联的对应的权益。
type TaobaoPromotionBenefitActivityRelationAPIRequest struct {
	model.Params
	// 活动关联权益请求参数
	_relationRequest *RelationActivityBenefitRequest
}

// NewTaobaoPromotionBenefitActivityRelationRequest 初始化TaobaoPromotionBenefitActivityRelationAPIRequest对象
func NewTaobaoPromotionBenefitActivityRelationRequest() *TaobaoPromotionBenefitActivityRelationAPIRequest {
	return &TaobaoPromotionBenefitActivityRelationAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionBenefitActivityRelationAPIRequest) Reset() {
	r._relationRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityRelationAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityRelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionBenefitActivityRelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationRequest is RelationRequest Setter
// 活动关联权益请求参数
func (r *TaobaoPromotionBenefitActivityRelationAPIRequest) SetRelationRequest(_relationRequest *RelationActivityBenefitRequest) error {
	r._relationRequest = _relationRequest
	r.Set("relation_request", _relationRequest)
	return nil
}

// GetRelationRequest RelationRequest Getter
func (r TaobaoPromotionBenefitActivityRelationAPIRequest) GetRelationRequest() *RelationActivityBenefitRequest {
	return r._relationRequest
}

var poolTaobaoPromotionBenefitActivityRelationAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionBenefitActivityRelationRequest()
	},
}

// GetTaobaoPromotionBenefitActivityRelationRequest 从 sync.Pool 获取 TaobaoPromotionBenefitActivityRelationAPIRequest
func GetTaobaoPromotionBenefitActivityRelationAPIRequest() *TaobaoPromotionBenefitActivityRelationAPIRequest {
	return poolTaobaoPromotionBenefitActivityRelationAPIRequest.Get().(*TaobaoPromotionBenefitActivityRelationAPIRequest)
}

// ReleaseTaobaoPromotionBenefitActivityRelationAPIRequest 将 TaobaoPromotionBenefitActivityRelationAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionBenefitActivityRelationAPIRequest(v *TaobaoPromotionBenefitActivityRelationAPIRequest) {
	v.Reset()
	poolTaobaoPromotionBenefitActivityRelationAPIRequest.Put(v)
}
