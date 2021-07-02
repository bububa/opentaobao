package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponEndactivityAPIRequest 结束优惠券活动 API请求
// alibaba.wdk.marketing.coupon.endactivity
//
// 结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
type AlibabaWdkMarketingCouponEndactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabaWdkMarketingCouponEndactivityRequest 初始化AlibabaWdkMarketingCouponEndactivityAPIRequest对象
func NewAlibabaWdkMarketingCouponEndactivityRequest() *AlibabaWdkMarketingCouponEndactivityAPIRequest {
	return &AlibabaWdkMarketingCouponEndactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponEndactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.endactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponEndactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingCouponEndactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaWdkMarketingCouponEndactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
