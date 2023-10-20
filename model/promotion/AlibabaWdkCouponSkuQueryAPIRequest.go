package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcouponskuqueryAPIRequest 优惠券商品查询 API请求
// alibaba.wdk.coupon.sku.query
//
// 优惠券商品查询
type AlibabawdkcouponskuqueryAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest
}

// NewAlibabawdkcouponskuqueryRequest 初始化AlibabawdkcouponskuqueryAPIRequest对象
func NewAlibabawdkcouponskuqueryRequest() *AlibabawdkcouponskuqueryAPIRequest {
	return &AlibabawdkcouponskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcouponskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcouponskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcouponskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateItemQueryRequest is ParamCouponTemplateItemQueryRequest Setter
// 请求
func (r *AlibabawdkcouponskuqueryAPIRequest) SetParamCouponTemplateItemQueryRequest(_paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest) error {
	r._paramCouponTemplateItemQueryRequest = _paramCouponTemplateItemQueryRequest
	r.Set("param_coupon_template_item_query_request", _paramCouponTemplateItemQueryRequest)
	return nil
}

// GetParamCouponTemplateItemQueryRequest ParamCouponTemplateItemQueryRequest Getter
func (r AlibabawdkcouponskuqueryAPIRequest) GetParamCouponTemplateItemQueryRequest() *CouponTemplateItemQueryRequest {
	return r._paramCouponTemplateItemQueryRequest
}
