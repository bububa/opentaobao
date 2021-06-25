package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取优惠券信息 APIRequest
alibaba.data.coupon.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaDataCouponGetRequest struct {
    model.Params

    // 客户端鉴权虚拟api使用
    unNamed   string 

}

func NewAlibabaDataCouponGetRequest() *AlibabaDataCouponGetRequest{
    return &AlibabaDataCouponGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDataCouponGetRequest) GetApiMethodName() string {
    return "alibaba.data.coupon.get"
}

func (r AlibabaDataCouponGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDataCouponGetRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

func (r AlibabaDataCouponGetRequest) GetUnNamed() string {
    return r.unNamed
}

