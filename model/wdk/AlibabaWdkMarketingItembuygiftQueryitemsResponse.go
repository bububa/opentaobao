package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询买赠活动下的商品 APIResponse
alibaba.wdk.marketing.itembuygift.queryitems

查询买赠活动下的商品
*/
type AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItembuygiftQueryitemsResponse `json:"alibaba_wdk_marketing_itembuygift_queryitems_response,omitempty"`
}

type AlibabaWdkMarketingItembuygiftQueryitemsResponse struct {

    // 查询返回结果
    Result   *MarketPageResult `json:"result,omitempty"`

}
