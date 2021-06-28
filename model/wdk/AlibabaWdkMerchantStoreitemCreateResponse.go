package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新建门店商品 APIResponse
alibaba.wdk.merchant.storeitem.create

新建门店商品
*/
type AlibabaWdkMerchantStoreitemCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMerchantStoreitemCreateResponse `json:"alibaba_wdk_merchant_storeitem_create_response,omitempty"` 
    AlibabaWdkMerchantStoreitemCreateResponse
}

/* model for simplify = false
type AlibabaWdkMerchantStoreitemCreateResponse struct {

    // success
    
    Suc   bool `json:"suc,omitempty"`
    

    // errorCode
    
    Errorcode   string `json:"errorcode,omitempty"`
    

    // errorDesc
    
    Errordesc   string `json:"errordesc,omitempty"`
    

}
*/

type AlibabaWdkMerchantStoreitemCreateResponse struct {

    // success
    Suc   bool `json:"suc,omitempty"`

    // errorCode
    Errorcode   string `json:"errorcode,omitempty"`

    // errorDesc
    Errordesc   string `json:"errordesc,omitempty"`

}
