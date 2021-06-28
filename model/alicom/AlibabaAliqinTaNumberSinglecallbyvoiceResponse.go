package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据号码tts单呼 APIResponse
alibaba.aliqin.ta.number.singlecallbyvoice

根据号码语音单呼
*/
type AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinTaNumberSinglecallbyvoiceResponse `json:"alibaba_aliqin_ta_number_singlecallbyvoice_response,omitempty"` 
    AlibabaAliqinTaNumberSinglecallbyvoiceResponse
}

/* model for simplify = false
type AlibabaAliqinTaNumberSinglecallbyvoiceResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaAliqinTaNumberSinglecallbyvoiceResult  *AlibabaAliqinTaNumberSinglecallbyvoiceResult `json:"alibaba_aliqin_ta_number_singlecallbyvoice_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinTaNumberSinglecallbyvoiceResponse struct {

    // 接口返回model
    Result   *AlibabaAliqinTaNumberSinglecallbyvoiceResult `json:"result,omitempty"`

}
