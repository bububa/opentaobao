package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品池排除商品【品类优惠使用】 APIResponse
alibaba.wdk.marketing.itempool.excludeskucode

品类优惠新增排除池
*/
type AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItempoolExcludeskucodeResponse `json:"alibaba_wdk_marketing_itempool_excludeskucode_response,omitempty"` 
    AlibabaWdkMarketingItempoolExcludeskucodeResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItempoolExcludeskucodeResponse struct {

    // 商品报名活动的返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItempoolExcludeskucodeResponse struct {

    // 商品报名活动的返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
