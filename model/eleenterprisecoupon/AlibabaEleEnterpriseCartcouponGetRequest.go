package eleenterprisecoupon

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取下单可用的优惠券 APIRequest
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

func NewAlibabaEleEnterpriseCartcouponGetRequest() *AlibabaEleEnterpriseCartcouponGetRequest{
    return &AlibabaEleEnterpriseCartcouponGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseCartcouponGetRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.cartcoupon.get"
}

func (r AlibabaEleEnterpriseCartcouponGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseCartcouponGetRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaEleEnterpriseCartcouponGetRequest) GetPhone() string {
    return r.phone
}

func (r *AlibabaEleEnterpriseCartcouponGetRequest) SetCartId(cartId string) error {
    r.cartId = cartId
    r.Set("cart_id", cartId)
    return nil
}

func (r AlibabaEleEnterpriseCartcouponGetRequest) GetCartId() string {
    return r.cartId
}

