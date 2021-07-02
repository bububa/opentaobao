package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponCreateactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 创建优惠券活动请求入参
func (r *AlibabaWdkMarketingCouponCreateactivityAPIRequest) SetParam(_param *CouponActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaWdkMarketingCouponCreateactivityAPIRequest) GetParam() *CouponActivity {
	return r._param
}
