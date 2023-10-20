package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcoupontemplatequeryAPIRequest 优惠券模版查询 API请求
// alibaba.wdk.coupon.template.query
//
// 优惠券模版查询
type AlibabawdkcoupontemplatequeryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramCouponTemplateQueryRequest *CouponTemplateQueryRequest
}

// NewAlibabawdkcoupontemplatequeryRequest 初始化AlibabawdkcoupontemplatequeryAPIRequest对象
func NewAlibabawdkcoupontemplatequeryRequest() *AlibabawdkcoupontemplatequeryAPIRequest {
	return &AlibabawdkcoupontemplatequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcoupontemplatequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcoupontemplatequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcoupontemplatequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCouponTemplateQueryRequest is ParamCouponTemplateQueryRequest Setter
// 系统自动生成
func (r *AlibabawdkcoupontemplatequeryAPIRequest) SetParamCouponTemplateQueryRequest(_paramCouponTemplateQueryRequest *CouponTemplateQueryRequest) error {
	r._paramCouponTemplateQueryRequest = _paramCouponTemplateQueryRequest
	r.Set("param_coupon_template_query_request", _paramCouponTemplateQueryRequest)
	return nil
}

// GetParamCouponTemplateQueryRequest ParamCouponTemplateQueryRequest Getter
func (r AlibabawdkcoupontemplatequeryAPIRequest) GetParamCouponTemplateQueryRequest() *CouponTemplateQueryRequest {
	return r._paramCouponTemplateQueryRequest
}
