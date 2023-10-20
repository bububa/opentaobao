package eleenterprisecoupon

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseCartcouponGetAPIRequest) Reset() {
	r._phone = ""
	r._cartId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.cartcoupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *AlibabaEleEnterpriseCartcouponGetAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetPhone() string {
	return r._phone
}

// SetCartId is CartId Setter
// 购物车id
func (r *AlibabaEleEnterpriseCartcouponGetAPIRequest) SetCartId(_cartId string) error {
	r._cartId = _cartId
	r.Set("cart_id", _cartId)
	return nil
}

// GetCartId CartId Getter
func (r AlibabaEleEnterpriseCartcouponGetAPIRequest) GetCartId() string {
	return r._cartId
}

var poolAlibabaEleEnterpriseCartcouponGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseCartcouponGetRequest()
	},
}

// GetAlibabaEleEnterpriseCartcouponGetRequest 从 sync.Pool 获取 AlibabaEleEnterpriseCartcouponGetAPIRequest
func GetAlibabaEleEnterpriseCartcouponGetAPIRequest() *AlibabaEleEnterpriseCartcouponGetAPIRequest {
	return poolAlibabaEleEnterpriseCartcouponGetAPIRequest.Get().(*AlibabaEleEnterpriseCartcouponGetAPIRequest)
}

// ReleaseAlibabaEleEnterpriseCartcouponGetAPIRequest 将 AlibabaEleEnterpriseCartcouponGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseCartcouponGetAPIRequest(v *AlibabaEleEnterpriseCartcouponGetAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseCartcouponGetAPIRequest.Put(v)
}
