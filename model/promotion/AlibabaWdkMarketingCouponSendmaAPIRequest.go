package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponSendmaAPIRequest 发放匿名码 API请求
// alibaba.wdk.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
type AlibabaWdkMarketingCouponSendmaAPIRequest struct {
	model.Params
	// 发放匿名码入参
	_param0 *CommonActivityParam
}

// NewAlibabaWdkMarketingCouponSendmaRequest 初始化AlibabaWdkMarketingCouponSendmaAPIRequest对象
func NewAlibabaWdkMarketingCouponSendmaRequest() *AlibabaWdkMarketingCouponSendmaAPIRequest {
	return &AlibabaWdkMarketingCouponSendmaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponSendmaAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.coupon.sendma"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponSendmaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingCouponSendmaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发放匿名码入参
func (r *AlibabaWdkMarketingCouponSendmaAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingCouponSendmaAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
