package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询优惠券活动 APIResponse
alibaba.wdk.marketing.coupon.queryactivity

查询优惠券活动
*/
type AlibabaWdkMarketingCouponQueryactivityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingCouponQueryactivityResponse `json:"alibaba_wdk_marketing_coupon_queryactivity_response,omitempty"` 
    AlibabaWdkMarketingCouponQueryactivityResponse
}

/* model for simplify = false
type AlibabaWdkMarketingCouponQueryactivityResponse struct {

    // 查询返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingCouponQueryactivityResponse struct {

    // 查询返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
