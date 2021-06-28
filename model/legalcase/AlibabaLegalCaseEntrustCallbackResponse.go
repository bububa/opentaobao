package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
委托回调接口 APIResponse
alibaba.legal.case.entrust.callback

委托回调接口
*/
type AlibabaLegalCaseEntrustCallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_legal_case_entrust_callback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // void
    
    Content   string `json:"content,omitempty" xml:"