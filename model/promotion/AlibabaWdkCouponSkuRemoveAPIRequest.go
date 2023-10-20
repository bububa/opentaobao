package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcouponskuremoveAPIRequest 优惠券商品删除 API请求
// alibaba.wdk.coupon.sku.remove
//
// 优惠券商品删除
type AlibabawdkcouponskuremoveAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemRequest *CouponTemplateItemRequest
}

// NewAlibabawdkcouponskuremoveRequest 初始化AlibabawdkcouponskuremoveAPIRequest对象
func NewAlibabawdkcouponskuremoveRequest() *AlibabawdkcouponskuremoveAPIRequest {
	return &AlibabawdkcouponskuremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcouponskuremoveAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.sku.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcouponskuremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcouponskuremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateItemRequest is ParamCouponTemplateItemRequest Setter
// 请求
func (r *AlibabawdkcouponskuremoveAPIRequest) SetParamCouponTemplateItemRequest(_paramCouponTemplateItemRequest *CouponTemplateItemRequest) error {
	r._paramCouponTemplateItemRequest = _paramCouponTemplateItemRequest
	r.Set("param_coupon_template_item_request", _paramCouponTemplateItemRequest)
	return nil
}

// GetParamCouponTemplateItemRequest ParamCouponTemplateItemRequest Getter
func (r AlibabawdkcouponskuremoveAPIRequest) GetParamCouponTemplateItemRequest() *CouponTemplateItemRequest {
	return r._paramCouponTemplateItemRequest
}
