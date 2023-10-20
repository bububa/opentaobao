package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingCouponQueryactivityAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingCouponQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询优惠券活动入参
func (r *AlibabaWdkMarketingCouponQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingCouponQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}

var poolAlibabaWdkMarketingCouponQueryactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingCouponQueryactivityRequest()
	},
}

// GetAlibabaWdkMarketingCouponQueryactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingCouponQueryactivityAPIRequest
func GetAlibabaWdkMarketingCouponQueryactivityAPIRequest() *AlibabaWdkMarketingCouponQueryactivityAPIRequest {
	return poolAlibabaWdkMarketingCouponQueryactivityAPIRequest.Get().(*AlibabaWdkMarketingCouponQueryactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingCouponQueryactivityAPIRequest 将 AlibabaWdkMarketingCouponQueryactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingCouponQueryactivityAPIRequest(v *AlibabaWdkMarketingCouponQueryactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingCouponQueryactivityAPIRequest.Put(v)
}
