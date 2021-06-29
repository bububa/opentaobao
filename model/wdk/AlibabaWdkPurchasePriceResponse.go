package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rt回传采购价 APIResponse
alibaba.wdk.purchase.price

猫超共享库存项目-rt回传采购价
*/
type AlibabaWdkPurchasePriceAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPurchasePriceResponse
}

type AlibabaWdkPurchasePriceResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_purchase_price_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // SYSTEM ERROR
    
    ReturnMsg   string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`

    
    // ERROR
    
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`

    
    // true
    
    ReturnSuccess   bool `json:"return_success,omitempty" xml:"return_success,omitempty"`

    
}
