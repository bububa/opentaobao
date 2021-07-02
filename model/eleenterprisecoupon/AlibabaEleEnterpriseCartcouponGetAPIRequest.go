package eleenterprisecoupon

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseCartcouponGetAPIRequest 获取下单可用的优惠券 API请求
// alibaba.ele.enterprise.cartcoupon.get
//
// 获取下单可用的优惠券
type AlibabaEleEnterpriseCartcouponGetAPIRequest struct {
	model.Params
	// 手机号
	_phone string
	// 购物车id
	_cartId string
}

// NewAlibabaEleEnterpriseCartcouponGetRequest 初始化AlibabaEleEnterpriseCartcouponGetAPIRequest对象
func NewAlibabaEleEnterpriseCartcouponGetRequest() *AlibabaEleEnterpriseCartcouponGetAPIRequest {
	return &AlibabaEleEnterpriseCartcouponGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.cartcoupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Phone Setter
// 手机号
func (r *AlibabaEleEnterpriseCartcouponGetAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// Get Phone Getter
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetPhone() string {
	return r._phone
}

// Set is CartId Setter
// 购物车id
func (r *AlibabaEleEnterpriseCartcouponGetAPIRequest) SetCartId(_cartId string) error {
	r._cartId = _cartId
	r.Set("cart_id", _cartId)
	return nil
}

// Get CartId Getter
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetCartId() string {
	return r._cartId
}
