package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingcouponqueryactivityAPIRequest 查询优惠券活动 API请求
// alibaba.wdk.marketing.coupon.queryactivity
//
// 查询优惠券活动
type AlibabawdkmarketingcouponqueryactivityAPIRequest struct {
	model.Params
	// 查询优惠券活动入参
	_param0 *CommonActivityParam
}

// NewAlibabawdkmarketingcouponqueryactivityRequest 初始化AlibabawdkmarketingcouponqueryactivityAPIRequest对象
func NewAlibabawdkmarketingcouponqueryactivityRequest() *AlibabawdkmarketingcouponqueryactivityAPIRequest {
	return &AlibabawdkmarketingcouponqueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingcouponqueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingcouponqueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingcouponqueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询优惠券活动入参
func (r *AlibabawdkmarketingcouponqueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingcouponqueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
