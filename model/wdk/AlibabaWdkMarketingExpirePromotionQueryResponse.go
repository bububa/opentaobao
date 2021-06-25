package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠查询 APIResponse
alibaba.wdk.marketing.expire.promotion.query

短保优惠查询
*/
type AlibabaWdkMarketingExpirePromotionQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingExpirePromotionQueryResponse `json:"alibaba_wdk_marketing_expire_promotion_query_response,omitempty"`
}

type AlibabaWdkMarketingExpirePromotionQueryResponse struct {

    // result
    Result   *MarketResult `json:"result,omitempty"`

}
