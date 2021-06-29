package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券校验 API请求
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

// 初始化AlibabaSellerCouponAuthVerifyRequest对象
func NewAlibabaSellerCouponAuthVerifyRequest() *AlibabaSellerCouponAuthVerifyRequest{
    return &AlibabaSellerCouponAuthVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSellerCouponAuthVerifyRequest) GetApiMethodName() string {
    return "alibaba.seller.coupon.auth.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSellerCouponAuthVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceCode Setter
// 服务代码
func (r *AlibabaSellerCouponAuthVerifyRequest) SetServiceCode(serviceCode string) error {
    r.serviceCode = serviceCode
    r.Set("service_code", serviceCode)
    return nil
}

// ServiceCode Getter
func (r AlibabaSellerCouponAuthVerifyRequest) GetServiceCode() string {
    return r.serviceCode
}
// CouponSeqNumber Setter
// 卡券验证码
func (r *AlibabaSellerCouponAuthVerifyRequest) SetCouponSeqNumber(couponSeqNumber string) error {
    r.couponSeqNumber = couponSeqNumber
    r.Set("coupon_seq_number", couponSeqNumber)
    return nil
}

// CouponSeqNumber Getter
func (r AlibabaSellerCouponAuthVerifyRequest) GetCouponSeqNumber() string {
    return r.couponSeqNumber
}
