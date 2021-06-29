package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券校验 APIRequest
alibaba.seller.coupon.auth.verify

优惠券校验
*/
type AlibabaSellerCouponAuthVerifyRequest struct {
    model.Params

    // 服务代码
    serviceCode   string 

    // 卡券验证码
    couponSeqNumber   string 

}

func NewAlibabaSellerCouponAuthVerifyRequest() *AlibabaSellerCouponAuthVerifyRequest{
    return &AlibabaSellerCouponAuthVerifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSellerCouponAuthVerifyRequest) GetApiMethodName() string {
    return "alibaba.seller.coupon.auth.verify"
}

func (r AlibabaSellerCouponAuthVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSellerCouponAuthVerifyRequest) SetServiceCode(serviceCode string) error {
    r.serviceCode = serviceCode
    r.Set("service_code", serviceCode)
    return nil
}

func (r AlibabaSellerCouponAuthVerifyRequest) GetServiceCode() string {
    return r.serviceCode
}

func (r *AlibabaSellerCouponAuthVerifyRequest) SetCouponSeqNumber(couponSeqNumber string) error {
    r.couponSeqNumber = couponSeqNumber
    r.Set("coupon_seq_number", couponSeqNumber)
    return nil
}

func (r AlibabaSellerCouponAuthVerifyRequest) GetCouponSeqNumber() string {
    return r.couponSeqNumber
}

