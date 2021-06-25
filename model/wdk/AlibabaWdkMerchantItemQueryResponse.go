package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家商品查询 APIResponse
alibaba.wdk.merchant.item.query

商家商品查询
*/
type AlibabaWdkMerchantItemQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMerchantItemQueryResponse `json:"alibaba_wdk_merchant_item_query_response,omitempty"`
}

type AlibabaWdkMerchantItemQueryResponse struct {

    // result
    Result   *AlibabaWdkMerchantItemQueryResult `json:"result,omitempty"`

}
