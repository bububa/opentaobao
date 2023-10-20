package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingcouponendactivityAPIRequest 结束优惠券活动 API请求
// alibaba.wdk.marketing.coupon.endactivity
//
// 结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
type AlibabawdkmarketingcouponendactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabawdkmarketingcouponendactivityRequest 初始化AlibabawdkmarketingcouponendactivityAPIRequest对象
func NewAlibabawdkmarketingcouponendactivityRequest() *AlibabawdkmarketingcouponendactivityAPIRequest {
	return &AlibabawdkmarketingcouponendactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingcouponendactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.endactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingcouponendactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingcouponendactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabawdkmarketingcouponendactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingcouponendactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
