package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品池新增商品 APIResponse
alibaba.wdk.marketing.itempool.item.add.async

新分组模型下新增商品
*/
type AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolItemAddAsyncResponse `json:"alibaba_wdk_marketing_itempool_item_add_async_response,omitempty"`
}

type AlibabaWdkMarketingItempoolItemAddAsyncResponse struct {

    // 结果信息
    Result   *MarketResult `json:"result,omitempty"`

}
