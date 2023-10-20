package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponCreateactivityAPIRequest 优惠券活动创建 API请求
// alibaba.wdk.marketing.coupon.createactivity
//
// 添加优惠券活动
type AlibabaWdkMarketingCouponCreateactivityAPIRequest struct {
	model.Params
	// 创建优惠券活动请求入参
	_param *CouponActivity
}

// NewAlibabaWdkMarketingCouponCreateactivityRequest 初始化AlibabaWdkMarketingCouponCreateactivityAPIRequest对象
func NewAlibabaWdkMarketingCouponCreateactivityRequest() *AlibabaWdkMarketingCouponCreateactivityAPIRequest {
	return &AlibabaWdkMarketingCouponCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingCouponCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingCouponCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建优惠券活动请求入参
func (r *AlibabaWdkMarketingCouponCreateactivityAPIRequest) SetParam(_param *CouponActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingCouponCreateactivityAPIRequest) GetParam() *CouponActivity {
	return r._param
}

var poolAlibabaWdkMarketingCouponCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingCouponCreateactivityRequest()
	},
}

// GetAlibabaWdkMarketingCouponCreateactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingCouponCreateactivityAPIRequest
func GetAlibabaWdkMarketingCouponCreateactivityAPIRequest() *AlibabaWdkMarketingCouponCreateactivityAPIRequest {
	return poolAlibabaWdkMarketingCouponCreateactivityAPIRequest.Get().(*AlibabaWdkMarketingCouponCreateactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingCouponCreateactivityAPIRequest 将 AlibabaWdkMarketingCouponCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingCouponCreateactivityAPIRequest(v *AlibabaWdkMarketingCouponCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingCouponCreateactivityAPIRequest.Put(v)
}
