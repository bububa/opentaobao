package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新建商品草稿 APIResponse
alibaba.wdk.merchant.item.createdraft

新建商品草稿erp接口
*/
type AlibabaWdkMerchantItemCreatedraftAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMerchantItemCreatedraftResponse `json:"alibaba_wdk_merchant_item_createdraft_response,omitempty"` 
    AlibabaWdkMerchantItemCreatedraftResponse
}

/* model for simplify = false
type AlibabaWdkMerchantItemCreatedraftResponse struct {

    // result
    
    Result  *struct {
        AlibabaWdkMerchantItemCreatedraftResult  *AlibabaWdkMerchantItemCreatedraftResult `json:"alibaba_wdk_merchant_item_createdraft_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMerchantItemCreatedraftResponse struct {

    // result
    Result   *AlibabaWdkMerchantItemCreatedraftResult `json:"result,omitempty"`

}
