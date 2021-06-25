package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量发布买赠商品 APIResponse
alibaba.wdk.marketing.buygift.item.add.async

批量发布买赠商品
*/
type AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingBuygiftItemAddAsyncResponse `json:"alibaba_wdk_marketing_buygift_item_add_async_response,omitempty"`
}

type AlibabaWdkMarketingBuygiftItemAddAsyncResponse struct {

    // 结果信息
    Result   *MarketResult `json:"result,omitempty"`

}
