package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateTerminateAPIRequest 优惠券模版终止 API请求
// alibaba.wdk.coupon.template.terminate
//
// 优惠券模版终止
type AlibabaWdkCouponTemplateTerminateAPIRequest struct {
	model.Params
	// 参数
	_paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest
}

// NewAlibabaWdkCouponTemplateTerminateRequest 初始化AlibabaWdkCouponTemplateTerminateAPIRequest对象
func NewAlibabaWdkCouponTemplateTerminateRequest() *AlibabaWdkCouponTemplateTerminateAPIRequest {
	return &AlibabaWdkCouponTemplateTerminateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateTerminateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.terminate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateTerminateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamCouponTemplateTerminateRequest is ParamCouponTemplateTerminateRequest Setter
// 参数
func (r *AlibabaWdkCouponTemplateTerminateAPIRequest) SetParamCouponTemplateTerminateRequest(_paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest) error {
	r._paramCouponTemplateTerminateRequest = _paramCouponTemplateTerminateRequest
	r.Set("param_coupon_template_terminate_request", _paramCouponTemplateTerminateRequest)
	return nil
}

// GetParamCouponTemplateTerminateRequest ParamCouponTemplateTerminateRequest Getter
func (r AlibabaWdkCouponTemplateTerminateAPIRequest) GetParamCouponTemplateTerminateRequest() *CouponTemplateTerminateRequest {
	return r._paramCouponTemplateTerminateRequest
}
