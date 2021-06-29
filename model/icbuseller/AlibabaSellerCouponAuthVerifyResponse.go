package icbuseller

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券校验 APIResponse
alibaba.seller.coupon.auth.verify

优惠券校验
*/
type AlibabaSellerCouponAuthVerifyAPIResponse struct {
    model.CommonResponse
    AlibabaSellerCouponAuthVerifyResponse
}

type AlibabaSellerCouponAuthVerifyResponse struct {
    XMLName xml.Name `xml:"alibaba_seller_coupon_auth_verify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取是否验证成功
    
    Result   *AlibabaSellerCouponAuthVerifyResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
