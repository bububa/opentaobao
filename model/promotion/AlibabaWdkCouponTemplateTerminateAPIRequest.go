package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcoupontemplateterminateAPIRequest 优惠券模版终止 API请求
// alibaba.wdk.coupon.template.terminate
//
// 优惠券模版终止
type AlibabawdkcoupontemplateterminateAPIRequest struct {
	model.Params
	// 参数
	_paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest
}

// NewAlibabawdkcoupontemplateterminateRequest 初始化AlibabawdkcoupontemplateterminateAPIRequest对象
func NewAlibabawdkcoupontemplateterminateRequest() *AlibabawdkcoupontemplateterminateAPIRequest {
	return &AlibabawdkcoupontemplateterminateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcoupontemplateterminateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.terminate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcoupontemplateterminateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcoupontemplateterminateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateTerminateRequest is ParamCouponTemplateTerminateRequest Setter
// 参数
func (r *AlibabawdkcoupontemplateterminateAPIRequest) SetParamCouponTemplateTerminateRequest(_paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest) error {
	r._paramCouponTemplateTerminateRequest = _paramCouponTemplateTerminateRequest
	r.Set("param_coupon_template_terminate_request", _paramCouponTemplateTerminateRequest)
	return nil
}

// GetParamCouponTemplateTerminateRequest ParamCouponTemplateTerminateRequest Getter
func (r AlibabawdkcoupontemplateterminateAPIRequest) GetParamCouponTemplateTerminateRequest() *CouponTemplateTerminateRequest {
	return r._paramCouponTemplateTerminateRequest
}
