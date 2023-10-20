package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcouponskuaddAPIRequest 优惠券商品增加 API请求
// alibaba.wdk.coupon.sku.add
//
// 优惠券商品增加
type AlibabawdkcouponskuaddAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemRequest *CouponTemplateItemRequest
}

// NewAlibabawdkcouponskuaddRequest 初始化AlibabawdkcouponskuaddAPIRequest对象
func NewAlibabawdkcouponskuaddRequest() *AlibabawdkcouponskuaddAPIRequest {
	return &AlibabawdkcouponskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcouponskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcouponskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcouponskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateItemRequest is ParamCouponTemplateItemRequest Setter
// 请求
func (r *AlibabawdkcouponskuaddAPIRequest) SetParamCouponTemplateItemRequest(_paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
	r._paramCouponTemplateItemRequest = _paramCouponTemplateItemRequest
	r.Set("param_coupon_template_item_request", _paramCouponTemplateItemRequest)
	return nil
}

// GetParamCouponTemplateItemRequest ParamCouponTemplateItemRequest Getter
func (r AlibabawdkcouponskuaddAPIRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
	return r._paramCouponTemplateItemRequest
}
