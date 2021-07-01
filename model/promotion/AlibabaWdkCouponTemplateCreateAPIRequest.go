package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponTemplateCreateAPIRequest
优惠券模版创建 API请求
alibaba.wdk.coupon.template.create

开放给外部商家创建优惠券模版 */
type AlibabaWdkCouponTemplateCreateAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest
}

// NewAlibabaWdkCouponTemplateCreateRequest 初始化AlibabaWdkCouponTemplateCreateAPIRequest对象
func NewAlibabaWdkCouponTemplateCreateRequest() *AlibabaWdkCouponTemplateCreateAPIRequest {
	return &AlibabaWdkCouponTemplateCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamCouponTemplateOperateRequest Setter
// 请求
func (r *AlibabaWdkCouponTemplateCreateAPIRequest) SetParamCouponTemplateOperateRequest(_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
	r._paramCouponTemplateOperateRequest = _paramCouponTemplateOperateRequest
	r.Set("param_coupon_template_operate_request", _paramCouponTemplateOperateRequest)
	return nil
}

// Get ParamCouponTemplateOperateRequest Getter
func (r AlibabaWdkCouponTemplateCreateAPIRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
	return r._paramCouponTemplateOperateRequest
}
