package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加商品池里面的商品 APIResponse
alibaba.wdk.marketing.itempool.additem

增加商品池里面的商品
*/
type AlibabaWdkMarketingItempoolAdditemAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolAdditemResponse `json:"alibaba_wdk_marketing_itempool_additem_response,omitempty"`
}

type AlibabaWdkMarketingItempoolAdditemResponse struct {

    // 商品报名活动的返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
