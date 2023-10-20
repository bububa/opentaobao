package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingCouponEndactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponEndactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.endactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponEndactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingCouponEndactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingCouponEndactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingCouponEndactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}

var poolAlibabaWdkMarketingCouponEndactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingCouponEndactivityRequest()
	},
}

// GetAlibabaWdkMarketingCouponEndactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingCouponEndactivityAPIRequest
func GetAlibabaWdkMarketingCouponEndactivityAPIRequest() *AlibabaWdkMarketingCouponEndactivityAPIRequest {
	return poolAlibabaWdkMarketingCouponEndactivityAPIRequest.Get().(*AlibabaWdkMarketingCouponEndactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingCouponEndactivityAPIRequest 将 AlibabaWdkMarketingCouponEndactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingCouponEndactivityAPIRequest(v *AlibabaWdkMarketingCouponEndactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingCouponEndactivityAPIRequest.Put(v)
}
