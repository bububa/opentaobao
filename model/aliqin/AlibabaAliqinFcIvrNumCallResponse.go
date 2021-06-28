package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ivr呼叫 APIResponse
alibaba.aliqin.fc.ivr.num.call

ivr呼叫
*/
type AlibabaAliqinFcIvrNumCallAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcIvrNumCallResponse `json:"alibaba_aliqin_fc_ivr_num_call_response,omitempty"` 
    AlibabaAliqinFcIvrNumCallResponse
}

/* model for simplify = false
type AlibabaAliqinFcIvrNumCallResponse struct {

    // result
    
    Result  *struct {
        AlibabaAliqinFcIvrNumCallResult  *AlibabaAliqinFcIvrNumCallResult `json:"alibaba_aliqin_fc_ivr_num_call_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcIvrNumCallResponse struct {

    // result
    Result   *AlibabaAliqinFcIvrNumCallResult `json:"result,omitempty"`

}
