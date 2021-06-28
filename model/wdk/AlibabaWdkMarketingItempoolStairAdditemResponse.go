package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品池阶梯商品添加 APIResponse
alibaba.wdk.marketing.itempool.stair.additem

添加商品池阶梯商品
*/
type AlibabaWdkMarketingItempoolStairAdditemAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItempoolStairAdditemResponse `json:"alibaba_wdk_marketing_itempool_stair_additem_response,omitempty"` 
    AlibabaWdkMarketingItempoolStairAdditemResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItempoolStairAdditemResponse struct {

    // 添加商品返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItempoolStairAdditemResponse struct {

    // 添加商品返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
