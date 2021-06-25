package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
换购商品查询 APIResponse
alibaba.wdk.marketing.itempool.stair.queryitem

换购商品查询
*/
type AlibabaWdkMarketingItempoolStairQueryitemAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolStairQueryitemResponse `json:"alibaba_wdk_marketing_itempool_stair_queryitem_response,omitempty"`
}

type AlibabaWdkMarketingItempoolStairQueryitemResponse struct {

    // 查询结果
    Result   *MarketPageResult `json:"result,omitempty"`

}
