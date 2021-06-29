package eleenterprisecoupon

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户优惠券 APIRequest
alibaba.ele.enterprise.coupon.get

获取用户优惠券
*/
type AlibabaEleEnterpriseCouponGetRequest struct {
    model.Params

    // 手机号
    phone   string 

}

func NewAlibabaEleEnterpriseCouponGetRequest() *AlibabaEleEnterpriseCouponGetRequest{
    return &AlibabaEleEnterpriseCouponGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseCouponGetRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.coupon.get"
}

func (r AlibabaEleEnterpriseCouponGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseCouponGetRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaEleEnterpriseCouponGetRequest) GetPhone() string {
    return r.phone
}

