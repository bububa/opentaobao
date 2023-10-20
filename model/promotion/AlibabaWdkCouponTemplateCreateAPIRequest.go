package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcoupontemplatecreateAPIRequest 优惠券模版创建 API请求
// alibaba.wdk.coupon.template.create
//
// 开放给外部商家创建优惠券模版
type AlibabawdkcoupontemplatecreateAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest
}

// NewAlibabawdkcoupontemplatecreateRequest 初始化AlibabawdkcoupontemplatecreateAPIRequest对象
func NewAlibabawdkcoupontemplatecreateRequest() *AlibabawdkcoupontemplatecreateAPIRequest {
	return &AlibabawdkcoupontemplatecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcoupontemplatecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcoupontemplatecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcoupontemplatecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateOperateRequest is ParamCouponTemplateOperateRequest Setter
// 请求
func (r *AlibabawdkcoupontemplatecreateAPIRequest) SetParamCouponTemplateOperateRequest(_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
	r._paramCouponTemplateOperateRequest = _paramCouponTemplateOperateRequest
	r.Set("param_coupon_template_operate_request", _paramCouponTemplateOperateRequest)
	return nil
}

// GetParamCouponTemplateOperateRequest ParamCouponTemplateOperateRequest Getter
func (r AlibabawdkcoupontemplatecreateAPIRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
	return r._paramCouponTemplateOperateRequest
}
