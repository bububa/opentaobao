package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingCouponEndactivityAPIRequest 结束优惠券活动 API请求
// alibaba.hm.marketing.coupon.endactivity
//
// 结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
type AlibabaHmMarketingCouponEndactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabaHmMarketingCouponEndactivityRequest 初始化AlibabaHmMarketingCouponEndactivityAPIRequest对象
func NewAlibabaHmMarketingCouponEndactivityRequest() *AlibabaHmMarketingCouponEndactivityAPIRequest {
	return &AlibabaHmMarketingCouponEndactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingCouponEndactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingCouponEndactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.coupon.endactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingCouponEndactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingCouponEndactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaHmMarketingCouponEndactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingCouponEndactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}

var poolAlibabaHmMarketingCouponEndactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingCouponEndactivityRequest()
	},
}

// GetAlibabaHmMarketingCouponEndactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingCouponEndactivityAPIRequest
func GetAlibabaHmMarketingCouponEndactivityAPIRequest() *AlibabaHmMarketingCouponEndactivityAPIRequest {
	return poolAlibabaHmMarketingCouponEndactivityAPIRequest.Get().(*AlibabaHmMarketingCouponEndactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingCouponEndactivityAPIRequest 将 AlibabaHmMarketingCouponEndactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingCouponEndactivityAPIRequest(v *AlibabaHmMarketingCouponEndactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingCouponEndactivityAPIRequest.Put(v)
}
