package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
报名特价商品 APIResponse
alibaba.wdk.marketing.itemdiscount.additem

在商品特价活动中报名特价商品
*/
type AlibabaWdkMarketingItemdiscountAdditemAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItemdiscountAdditemResponse `json:"alibaba_wdk_marketing_itemdiscount_additem_response,omitempty"`
}

type AlibabaWdkMarketingItemdiscountAdditemResponse struct {

    // 商品报名活动的返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
