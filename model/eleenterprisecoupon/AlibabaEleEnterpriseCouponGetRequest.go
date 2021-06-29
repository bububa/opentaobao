package eleenterprisecoupon

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户优惠券 API请求
alibaba.ele.enterprise.coupon.get

获取用户优惠券
*/
type AlibabaEleEnterpriseCouponGetRequest struct {
    model.Params
    // 手机号
    _phone   string
}

// 初始化AlibabaEleEnterpriseCouponGetRequest对象
func NewAlibabaEleEnterpriseCouponGetRequest() *AlibabaEleEnterpriseCouponGetRequest{
    return &AlibabaEleEnterpriseCouponGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCouponGetRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.coupon.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCouponGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Phone Setter
// 手机号
func (r *AlibabaEleEnterpriseCouponGetRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCouponGetRequest) GetPhone() string {
    return r._phone
}
