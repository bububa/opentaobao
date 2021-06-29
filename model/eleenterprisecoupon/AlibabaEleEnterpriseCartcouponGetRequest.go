package eleenterprisecoupon

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取下单可用的优惠券 API请求
alibaba.ele.enterprise.cartcoupon.get

获取下单可用的优惠券
*/
type AlibabaEleEnterpriseCartcouponGetRequest struct {
    model.Params
    // 手机号
    _phone   string
    // 购物车id
    _cartId   string
}

// 初始化AlibabaEleEnterpriseCartcouponGetRequest对象
func NewAlibabaEleEnterpriseCartcouponGetRequest() *AlibabaEleEnterpriseCartcouponGetRequest{
    return &AlibabaEleEnterpriseCartcouponGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCartcouponGetRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.cartcoupon.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCartcouponGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Phone Setter
// 手机号
func (r *AlibabaEleEnterpriseCartcouponGetRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCartcouponGetRequest) GetPhone() string {
    return r._phone
}
// CartId Setter
// 购物车id
func (r *AlibabaEleEnterpriseCartcouponGetRequest) SetCartId(_cartId string) error {
    r._cartId = _cartId
    r.Set("cart_id", _cartId)
    return nil
}

// CartId Getter
func (r AlibabaEleEnterpriseCartcouponGetRequest) GetCartId() string {
    return r._cartId
}
