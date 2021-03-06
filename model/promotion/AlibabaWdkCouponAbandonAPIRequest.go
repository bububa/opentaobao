package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponAbandonAPIRequest 废券 API请求
// alibaba.wdk.coupon.abandon
//
// 优惠券废弃
type AlibabaWdkCouponAbandonAPIRequest struct {
	model.Params
	// 废券参数
	_paramWdkCouponAbandonParam *WdkCouponAbandonParam
}

// NewAlibabaWdkCouponAbandonRequest 初始化AlibabaWdkCouponAbandonAPIRequest对象
func NewAlibabaWdkCouponAbandonRequest() *AlibabaWdkCouponAbandonAPIRequest {
	return &AlibabaWdkCouponAbandonAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponAbandonAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.abandon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponAbandonAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamWdkCouponAbandonParam is ParamWdkCouponAbandonParam Setter
// 废券参数
func (r *AlibabaWdkCouponAbandonAPIRequest) SetParamWdkCouponAbandonParam(_paramWdkCouponAbandonParam *WdkCouponAbandonParam) error {
	r._paramWdkCouponAbandonParam = _paramWdkCouponAbandonParam
	r.Set("param_wdk_coupon_abandon_param", _paramWdkCouponAbandonParam)
	return nil
}

// GetParamWdkCouponAbandonParam ParamWdkCouponAbandonParam Getter
func (r AlibabaWdkCouponAbandonAPIRequest) GetParamWdkCouponAbandonParam() *WdkCouponAbandonParam {
	return r._paramWdkCouponAbandonParam
}
