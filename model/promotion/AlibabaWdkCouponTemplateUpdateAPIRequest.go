package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateUpdateAPIRequest 优惠券模版修改 API请求
// alibaba.wdk.coupon.template.update
//
// 优惠券模版修改
type AlibabaWdkCouponTemplateUpdateAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest
}

// NewAlibabaWdkCouponTemplateUpdateRequest 初始化AlibabaWdkCouponTemplateUpdateAPIRequest对象
func NewAlibabaWdkCouponTemplateUpdateRequest() *AlibabaWdkCouponTemplateUpdateAPIRequest {
	return &AlibabaWdkCouponTemplateUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamCouponTemplateOperateRequest Setter
// 请求
func (r *AlibabaWdkCouponTemplateUpdateAPIRequest) SetParamCouponTemplateOperateRequest(_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
	r._paramCouponTemplateOperateRequest = _paramCouponTemplateOperateRequest
	r.Set("param_coupon_template_operate_request", _paramCouponTemplateOperateRequest)
	return nil
}

// Get ParamCouponTemplateOperateRequest Getter
func (r AlibabaWdkCouponTemplateUpdateAPIRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
	return r._paramCouponTemplateOperateRequest
}
