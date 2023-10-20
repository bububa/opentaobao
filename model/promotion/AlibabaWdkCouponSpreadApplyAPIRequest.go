package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcouponspreadapplyAPIRequest 普通发券 API请求
// alibaba.wdk.coupon.spread.apply
//
// 优惠券发放
type AlibabawdkcouponspreadapplyAPIRequest struct {
	model.Params
	// 参数对象
	_paramWdkCouponApplyParam *WdkCouponApplyParam
}

// NewAlibabawdkcouponspreadapplyRequest 初始化AlibabawdkcouponspreadapplyAPIRequest对象
func NewAlibabawdkcouponspreadapplyRequest() *AlibabawdkcouponspreadapplyAPIRequest {
	return &AlibabawdkcouponspreadapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkcouponspreadapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.spread.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkcouponspreadapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkcouponspreadapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamWdkCouponApplyParam is ParamWdkCouponApplyParam Setter
// 参数对象
func (r *AlibabawdkcouponspreadapplyAPIRequest) SetParamWdkCouponApplyParam(_paramWdkCouponApplyParam *WdkCouponApplyParam) error {
	r._paramWdkCouponApplyParam = _paramWdkCouponApplyParam
	r.Set("param_wdk_coupon_apply_param", _paramWdkCouponApplyParam)
	return nil
}

// GetParamWdkCouponApplyParam ParamWdkCouponApplyParam Getter
func (r AlibabawdkcouponspreadapplyAPIRequest) GetParamWdkCouponApplyParam() *WdkCouponApplyParam {
	return r._paramWdkCouponApplyParam
}
