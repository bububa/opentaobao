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
    phone   string
    // 购物车id
    cartId   string
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
func (r *AlibabaEleEnterpriseCartcouponGetRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCartcouponGetRequest) GetPhone() string {
    return r.phone
}
// CartId Setter
// 购物车id
func (r *AlibabaEleEnterpriseCartcouponGetRequest) SetCartId(cartId string) error {
    r.cartId = cartId
    r.Set("cart_id", cartId)
    return nil
}

// CartId Getter
func (r AlibabaEleEnterpriseCartcouponGetRequest) GetCartId() string {
    return r.cartId
}
