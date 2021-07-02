package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponQueryactivityAPIRequest 查询优惠券活动 API请求
// alibaba.wdk.marketing.coupon.queryactivity
//
// 查询优惠券活动
type AlibabaWdkMarketingCouponQueryactivityAPIRequest struct {
	model.Params
	// 查询优惠券活动入参
	_param0 *CommonActivityParam
}

// NewAlibabaWdkMarketingCouponQueryactivityRequest 初始化AlibabaWdkMarketingCouponQueryactivityAPIRequest对象
func NewAlibabaWdkMarketingCouponQueryactivityRequest() *AlibabaWdkMarketingCouponQueryactivityAPIRequest {
	return &AlibabaWdkMarketingCouponQueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponQueryactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 查询优惠券活动入参
func (r *AlibabaWdkMarketingCouponQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingCouponQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
