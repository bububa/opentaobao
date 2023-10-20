package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcouponabandonAPIRequest 废券 API请求
// alibaba.wdk.coupon.abandon
//
// 优惠券废弃
type AlibabawdkcouponabandonAPIRequest struct {
	model.Params
	// 废券参数
	_paramWdkCouponAbandonParam *WdkCouponAbandonParam
}

// NewAlibabawdkcouponabandonRequest 初始化AlibabawdkcouponabandonAPIRequest对象
func NewAlibabawdkcouponabandonRequest() *AlibabawdkcouponabandonAPIRequest {
	return &AlibabawdkcouponabandonAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcouponabandonAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.abandon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcouponabandonAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcouponabandonAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWdkCouponAbandonParam is ParamWdkCouponAbandonParam Setter
// 废券参数
func (r *AlibabawdkcouponabandonAPIRequest) SetParamWdkCouponAbandonParam(_paramWdkCouponAbandonParam *WdkCouponAbandonParam) error {
	r._paramWdkCouponAbandonParam = _paramWdkCouponAbandonParam
	r.Set("param_wdk_coupon_abandon_param", _paramWdkCouponAbandonParam)
	return nil
}

// GetParamWdkCouponAbandonParam ParamWdkCouponAbandonParam Getter
func (r AlibabawdkcouponabandonAPIRequest) GetParamWdkCouponAbandonParam() *WdkCouponAbandonParam {
	return r._paramWdkCouponAbandonParam
}
