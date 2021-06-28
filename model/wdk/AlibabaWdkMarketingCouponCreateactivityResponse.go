package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券活动创建 APIResponse
alibaba.wdk.marketing.coupon.createactivity

添加优惠券活动
*/
type AlibabaWdkMarketingCouponCreateactivityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingCouponCreateactivityResponse `json:"alibaba_wdk_marketing_coupon_createactivity_response,omitempty"` 
    AlibabaWdkMarketingCouponCreateactivityResponse
}

/* model for simplify = false
type AlibabaWdkMarketingCouponCreateactivityResponse struct {

    // 创建优惠券活动返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingCouponCreateactivityResponse struct {

    // 创建优惠券活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
