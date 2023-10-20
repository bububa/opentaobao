package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcoupontemplateupdateAPIRequest 优惠券模版修改 API请求
// alibaba.wdk.coupon.template.update
//
// 优惠券模版修改
type AlibabawdkcoupontemplateupdateAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest
}

// NewAlibabawdkcoupontemplateupdateRequest 初始化AlibabawdkcoupontemplateupdateAPIRequest对象
func NewAlibabawdkcoupontemplateupdateRequest() *AlibabawdkcoupontemplateupdateAPIRequest {
	return &AlibabawdkcoupontemplateupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcoupontemplateupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcoupontemplateupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcoupontemplateupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateOperateRequest is ParamCouponTemplateOperateRequest Setter
// 请求
func (r *AlibabawdkcoupontemplateupdateAPIRequest) SetParamCouponTemplateOperateRequest(_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest) error {
	r._paramCouponTemplateOperateRequest = _paramCouponTemplateOperateRequest
	r.Set("param_coupon_template_operate_request", _paramCouponTemplateOperateRequest)
	return nil
}

// GetParamCouponTemplateOperateRequest ParamCouponTemplateOperateRequest Getter
func (r AlibabawdkcoupontemplateupdateAPIRequest) GetParamCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
	return r._paramCouponTemplateOperateRequest
}
