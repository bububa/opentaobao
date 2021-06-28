package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发放匿名码 APIResponse
alibaba.wdk.marketing.coupon.sendma

根据优惠券活动id打印单个匿名码
*/
type AlibabaWdkMarketingCouponSendmaAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingCouponSendmaResponse `json:"alibaba_wdk_marketing_coupon_sendma_response,omitempty"` 
    AlibabaWdkMarketingCouponSendmaResponse
}

/* model for simplify = false
type AlibabaWdkMarketingCouponSendmaResponse struct {

    // 发放匿名码返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingCouponSendmaResponse struct {

    // 发放匿名码返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
