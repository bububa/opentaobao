package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityDetailGetAPIRequest 活动关联的权益详情获取 API请求
// taobao.promotion.benefit.activity.detail.get
//
// 活动关联的权益详情获取
type TaobaoPromotionBenefitActivityDetailGetAPIRequest struct {
	model.Params
	// 查询活动关联权益详情的请求
	_queryRequest *ActivityRelationDetailRequest
}

// NewTaobaoPromotionBenefitActivityDetailGetRequest 初始化TaobaoPromotionBenefitActivityDetailGetAPIRequest对象
func NewTaobaoPromotionBenefitActivityDetailGetRequest() *TaobaoPromotionBenefitActivityDetailGetAPIRequest {
	return &TaobaoPromotionBenefitActivityDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityDetailGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryRequest is QueryRequest Setter
// 查询活动关联权益详情的请求
func (r *TaobaoPromotionBenefitActivityDetailGetAPIRequest) SetQueryRequest(_queryRequest *ActivityRelationDetailRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r TaobaoPromotionBenefitActivityDetailGetAPIRequest) GetQueryRequest() *ActivityRelationDetailRequest {
	return r._queryRequest
}
