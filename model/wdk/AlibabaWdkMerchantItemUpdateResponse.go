package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改商家商品 APIResponse
alibaba.wdk.merchant.item.update

修改商家商品
*/
type AlibabaWdkMerchantItemUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMerchantItemUpdateResponse `json:"alibaba_wdk_merchant_item_update_response,omitempty"` 
    AlibabaWdkMerchantItemUpdateResponse
}

/* model for simplify = false
type AlibabaWdkMerchantItemUpdateResponse struct {

    // result
    
    Result  *struct {
        AlibabaWdkMerchantItemUpdateResult  *AlibabaWdkMerchantItemUpdateResult `json:"alibaba_wdk_merchant_item_update_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMerchantItemUpdateResponse struct {

    // result
    Result   *AlibabaWdkMerchantItemUpdateResult `json:"result,omitempty"`

}
