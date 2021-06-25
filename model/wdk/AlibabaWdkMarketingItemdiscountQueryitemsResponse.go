package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询特价商品 APIResponse
alibaba.wdk.marketing.itemdiscount.queryitems

查询参加特价活动的商品优惠详情
*/
type AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItemdiscountQueryitemsResponse `json:"alibaba_wdk_marketing_itemdiscount_queryitems_response,omitempty"`
}

type AlibabaWdkMarketingItemdiscountQueryitemsResponse struct {

    // 查询返回结果
    Result   *MarketPageResult `json:"result,omitempty"`

}
