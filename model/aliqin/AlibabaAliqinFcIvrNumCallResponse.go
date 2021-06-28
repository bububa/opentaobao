package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ivr呼叫 APIResponse
alibaba.aliqin.fc.ivr.num.call

ivr呼叫
*/
type AlibabaAliqinFcIvrNumCallAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIvrNumCallResponse
}

type AlibabaAliqinFcIvrNumCallResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_ivr_num_call_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaAliqinFcIvrNumCallResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
