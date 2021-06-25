package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
rt回传采购价 APIResponse
alibaba.wdk.purchase.price

猫超共享库存项目-rt回传采购价
*/
type AlibabaWdkPurchasePriceAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkPurchasePriceResponse `json:"alibaba_wdk_purchase_price_response,omitempty"`
}

type AlibabaWdkPurchasePriceResponse struct {

    // SYSTEM ERROR
    ReturnMsg   string `json:"return_msg,omitempty"`

    // ERROR
    ReturnCode   string `json:"return_code,omitempty"`

    // true
    ReturnSuccess   bool `json:"return_success,omitempty"`

}
