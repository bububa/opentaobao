package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取优惠券信息 API请求
alibaba.data.coupon.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaDataCouponGetRequest struct {
    model.Params
    // 客户端鉴权虚拟api使用
    _unNamed   string
}

// 初始化AlibabaDataCouponGetRequest对象
func NewAlibabaDataCouponGetRequest() *AlibabaDataCouponGetRequest{
    return &AlibabaDataCouponGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDataCouponGetRequest) GetApiMethodName() string {
    return "alibaba.data.coupon.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDataCouponGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 客户端鉴权虚拟api使用
func (r *AlibabaDataCouponGetRequest) SetUnNamed(_unNamed string) error {
    r._unNamed = _unNamed
    r.Set("un_named", _unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaDataCouponGetRequest) GetUnNamed() string {
    return r._unNamed
}
