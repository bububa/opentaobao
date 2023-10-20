package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingCouponSendmaAPIRequest 发放匿名码 API请求
// alibaba.hm.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
type AlibabaHmMarketingCouponSendmaAPIRequest struct {
	model.Params
	// 发放匿名码入参
	_param0 *CommonActivityParam
}

// NewAlibabaHmMarketingCouponSendmaRequest 初始化AlibabaHmMarketingCouponSendmaAPIRequest对象
func NewAlibabaHmMarketingCouponSendmaRequest() *AlibabaHmMarketingCouponSendmaAPIRequest {
	return &AlibabaHmMarketingCouponSendmaAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingCouponSendmaAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingCouponSendmaAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.coupon.sendma"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingCouponSendmaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingCouponSendmaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发放匿名码入参
func (r *AlibabaHmMarketingCouponSendmaAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingCouponSendmaAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}

var poolAlibabaHmMarketingCouponSendmaAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingCouponSendmaRequest()
	},
}

// GetAlibabaHmMarketingCouponSendmaRequest 从 sync.Pool 获取 AlibabaHmMarketingCouponSendmaAPIRequest
func GetAlibabaHmMarketingCouponSendmaAPIRequest() *AlibabaHmMarketingCouponSendmaAPIRequest {
	return poolAlibabaHmMarketingCouponSendmaAPIRequest.Get().(*AlibabaHmMarketingCouponSendmaAPIRequest)
}

// ReleaseAlibabaHmMarketingCouponSendmaAPIRequest 将 AlibabaHmMarketingCouponSendmaAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingCouponSendmaAPIRequest(v *AlibabaHmMarketingCouponSendmaAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingCouponSendmaAPIRequest.Put(v)
}
