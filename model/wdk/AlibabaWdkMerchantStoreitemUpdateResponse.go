package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改门店商品 APIResponse
alibaba.wdk.merchant.storeitem.update

修改门店商品
*/
type AlibabaWdkMerchantStoreitemUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMerchantStoreitemUpdateResponse `json:"alibaba_wdk_merchant_storeitem_update_response,omitempty"` 
    AlibabaWdkMerchantStoreitemUpdateResponse
}

/* model for simplify = false
type AlibabaWdkMerchantStoreitemUpdateResponse struct {

    // result
    
    Result  *struct {
        AlibabaWdkMerchantStoreitemUpdateResult  *AlibabaWdkMerchantStoreitemUpdateResult `json:"alibaba_wdk_merchant_storeitem_update_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMerchantStoreitemUpdateResponse struct {

    // result
    Result   *AlibabaWdkMerchantStoreitemUpdateResult `json:"result,omitempty"`

}
