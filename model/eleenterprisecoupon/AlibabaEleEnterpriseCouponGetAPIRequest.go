package eleenterprisecoupon

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterprisecoupongetAPIRequest 获取用户优惠券 API请求
// alibaba.ele.enterprise.coupon.get
//
// 获取用户优惠券
type AlibabaeleenterprisecoupongetAPIRequest struct {
	model.Params
	// 手机号
	_phone string
}

// NewAlibabaeleenterprisecoupongetRequest 初始化AlibabaeleenterprisecoupongetAPIRequest对象
func NewAlibabaeleenterprisecoupongetRequest() *AlibabaeleenterprisecoupongetAPIRequest {
	return &AlibabaeleenterprisecoupongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterprisecoupongetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.coupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterprisecoupongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterprisecoupongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *AlibabaeleenterprisecoupongetAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaeleenterprisecoupongetAPIRequest) GetPhone() string {
	return r._phone
}
