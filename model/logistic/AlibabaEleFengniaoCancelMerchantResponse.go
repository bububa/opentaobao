package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商户取消 APIResponse
alibaba.ele.fengniao.cancel.merchant

商户取消配送
*/
type AlibabaEleFengniaoCancelMerchantAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaEleFengniaoCancelMerchantResponse `json:"alibaba_ele_fengniao_cancel_merchant_response,omitempty"` 
    AlibabaEleFengniaoCancelMerchantResponse
}

/* model for simplify = false
type AlibabaEleFengniaoCancelMerchantResponse struct {

    // message
    
    Message   string `json:"message,omitempty"`
    

}
*/

type AlibabaEleFengniaoCancelMerchantResponse struct {

    // message
    Message   string `json:"message,omitempty"`

}
