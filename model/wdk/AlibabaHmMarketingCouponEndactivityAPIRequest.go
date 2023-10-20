package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingcouponendactivityAPIRequest 结束优惠券活动 API请求
// alibaba.hm.marketing.coupon.endactivity
//
// 结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
type AlibabahmmarketingcouponendactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabahmmarketingcouponendactivityRequest 初始化AlibabahmmarketingcouponendactivityAPIRequest对象
func NewAlibabahmmarketingcouponendactivityRequest() *AlibabahmmarketingcouponendactivityAPIRequest {
	return &AlibabahmmarketingcouponendactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingcouponendactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.coupon.endactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingcouponendactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingcouponendactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabahmmarketingcouponendactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingcouponendactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
