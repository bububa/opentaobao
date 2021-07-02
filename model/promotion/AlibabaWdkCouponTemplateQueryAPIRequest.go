package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateQueryAPIRequest 优惠券模版查询 API请求
// alibaba.wdk.coupon.template.query
//
// 优惠券模版查询
type AlibabaWdkCouponTemplateQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramCouponTemplateQueryRequest *CouponTemplateQueryRequest
}

// NewAlibabaWdkCouponTemplateQueryRequest 初始化AlibabaWdkCouponTemplateQueryAPIRequest对象
func NewAlibabaWdkCouponTemplateQueryRequest() *AlibabaWdkCouponTemplateQueryAPIRequest {
	return &AlibabaWdkCouponTemplateQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamCouponTemplateQueryRequest is ParamCouponTemplateQueryRequest Setter
// 系统自动生成
func (r *AlibabaWdkCouponTemplateQueryAPIRequest) SetParamCouponTemplateQueryRequest(_paramCouponTemplateQueryRequest *CouponTemplateQueryRequest) error {
	r._paramCouponTemplateQueryRequest = _paramCouponTemplateQueryRequest
	r.Set("param_coupon_template_query_request", _paramCouponTemplateQueryRequest)
	return nil
}

// GetParamCouponTemplateQueryRequest ParamCouponTemplateQueryRequest Getter
func (r AlibabaWdkCouponTemplateQueryAPIRequest) GetParamCouponTemplateQueryRequest() *CouponTemplateQueryRequest {
	return r._paramCouponTemplateQueryRequest
}
