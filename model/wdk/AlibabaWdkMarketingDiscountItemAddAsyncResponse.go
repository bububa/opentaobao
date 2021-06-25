package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
特价批量新增商品 APIResponse
alibaba.wdk.marketing.discount.item.add.async

新分组模型下新增商品
*/
type AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingDiscountItemAddAsyncResponse `json:"alibaba_wdk_marketing_discount_item_add_async_response,omitempty"`
}

type AlibabaWdkMarketingDiscountItemAddAsyncResponse struct {

    // 结果信息
    Result   *MarketResult `json:"result,omitempty"`

}
