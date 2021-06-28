package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询优惠券活动下的商品 APIResponse
alibaba.wdk.marketing.coupon.queryitems

查询优惠券活动下面的商品
*/
type AlibabaWdkMarketingCouponQueryitemsAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingCouponQueryitemsResponse `json:"alibaba_wdk_marketing_coupon_queryitems_response,omitempty"` 
    AlibabaWdkMarketingCouponQueryitemsResponse
}

/* model for simplify = false
type AlibabaWdkMarketingCouponQueryitemsResponse struct {

    // 查询返回结果
    
    Result  *struct {
        MarketPageResult  *MarketPageResult `json:"market_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingCouponQueryitemsResponse struct {

    // 查询返回结果
    Result   *MarketPageResult `json:"result,omitempty"`

}
