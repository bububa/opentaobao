package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
委托回调接口 APIResponse
alibaba.legal.case.entrust.callback

委托回调接口
*/
type AlibabaLegalCaseEntrustCallbackAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLegalCaseEntrustCallbackResponse `json:"alibaba_legal_case_entrust_callback_response,omitempty"` 
    AlibabaLegalCaseEntrustCallbackResponse
}

/* model for simplify = false
type AlibabaLegalCaseEntrustCallbackResponse struct {

    // void
    
    Content   string `json:"content,omitempty"`
    

    // error_code
    
    ECode   string `json:"e_code,omitempty"`
    

    // error_msg
    
    EMsg   string `json:"e_msg,omitempty"`
    

}
*/

type AlibabaLegalCaseEntrustCallbackResponse struct {

    // void
    Content   string `json:"content,omitempty"`

    // error_code
    ECode   string `json:"e_code,omitempty"`

    // error_msg
    EMsg   string `json:"e_msg,omitempty"`

}
