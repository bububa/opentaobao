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
type AlibabaEleEnterpriseCouponGetAPIRequest struct {
    model.Params
    // 手机号
    _phone   string
}

// 初始化AlibabaEleEnterpriseCouponGetAPIRequest对象
func NewAlibabaEleEnterpriseCouponGetRequest() *AlibabaEleEnterpriseCouponGetAPIRequest{
    return &AlibabaEleEnterpriseCouponGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCouponGetAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.coupon.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCouponGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Phone Setter
// 手机号
func (r *AlibabaEleEnterpriseCouponGetAPIRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCouponGetAPIRequest) GetPhone() string {
    return r._phone
}
