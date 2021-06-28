package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
结束优惠券活动 APIResponse
alibaba.wdk.marketing.coupon.endactivity

结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
*/
type AlibabaWdkMarketingCouponEndactivityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingCouponEndactivityResponse `json:"alibaba_wdk_marketing_coupon_endactivity_response,omitempty"` 
    AlibabaWdkMarketingCouponEndactivityResponse
}

/* model for simplify = false
type AlibabaWdkMarketingCouponEndactivityResponse struct {

    // 删除活动返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingCouponEndactivityResponse struct {

    // 删除活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
