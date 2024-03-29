package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingCouponCreateactivityAPIRequest 优惠券活动创建 API请求
// alibaba.hm.marketing.coupon.createactivity
//
// 添加优惠券活动
type AlibabaHmMarketingCouponCreateactivityAPIRequest struct {
	model.Params
	// 创建优惠券活动请求入参
	_param *CouponActivity
}

// NewAlibabaHmMarketingCouponCreateactivityRequest 初始化AlibabaHmMarketingCouponCreateactivityAPIRequest对象
func NewAlibabaHmMarketingCouponCreateactivityRequest() *AlibabaHmMarketingCouponCreateactivityAPIRequest {
	return &AlibabaHmMarketingCouponCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingCouponCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingCouponCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.coupon.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingCouponCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingCouponCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建优惠券活动请求入参
func (r *AlibabaHmMarketingCouponCreateactivityAPIRequest) SetParam(_param *CouponActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingCouponCreateactivityAPIRequest) GetParam() *CouponActivity {
	return r._param
}

var poolAlibabaHmMarketingCouponCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingCouponCreateactivityRequest()
	},
}

// GetAlibabaHmMarketingCouponCreateactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingCouponCreateactivityAPIRequest
func GetAlibabaHmMarketingCouponCreateactivityAPIRequest() *AlibabaHmMarketingCouponCreateactivityAPIRequest {
	return poolAlibabaHmMarketingCouponCreateactivityAPIRequest.Get().(*AlibabaHmMarketingCouponCreateactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingCouponCreateactivityAPIRequest 将 AlibabaHmMarketingCouponCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingCouponCreateactivityAPIRequest(v *AlibabaHmMarketingCouponCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingCouponCreateactivityAPIRequest.Put(v)
}
