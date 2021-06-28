package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据号码tts单呼 APIResponse
alibaba.aliqin.ta.number.singlecallbyvoice

根据号码语音单呼
*/
type AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_ta_number_singlecallbyvoice_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaAliqinTaNumberSinglecallbyvoiceResult `json:"result,omitempty" xml:"