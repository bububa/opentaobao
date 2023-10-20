package eleenterprisecoupon

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterprisecartcoupongetAPIRequest 获取下单可用的优惠券 API请求
// alibaba.ele.enterprise.cartcoupon.get
//
// 获取下单可用的优惠券
type AlibabaeleenterprisecartcoupongetAPIRequest struct {
	model.Params
	// 手机号
	_phone string
	// 购物车id
	_cartId string
}

// NewAlibabaeleenterprisecartcoupongetRequest 初始化AlibabaeleenterprisecartcoupongetAPIRequest对象
func NewAlibabaeleenterprisecartcoupongetRequest() *AlibabaeleenterprisecartcoupongetAPIRequest {
	return &AlibabaeleenterprisecartcoupongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterprisecartcoupongetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.cartcoupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterprisecartcoupongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterprisecartcoupongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *AlibabaeleenterprisecartcoupongetAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaeleenterprisecartcoupongetAPIRequest) GetPhone() string {
	return r._phone
}

// SetCartId is CartId Setter
// 购物车id
func (r *AlibabaeleenterprisecartcoupongetAPIRequest) SetCartId(_cartId string) error {
	r._cartId = _cartId
	r.Set("cart_id", _cartId)
	return nil
}

// GetCartId CartId Getter
func (r AlibabaeleenterprisecartcoupongetAPIRequest) GetCartId() string {
	return r._cartId
}
