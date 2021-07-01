package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
委托回调接口 API返回值 
alibaba.legal.case.entrust.callback

委托回调接口
*/
type AlibabaLegalCaseEntrustCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaLegalCaseEntrustCallbackAPIResponseModel
}

// 委托回调接口 成功返回结果
type AlibabaLegalCaseEntrustCallbackAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_legal_case_entrust_callback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // void
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    // error_code
    ECode   string `json:"e_code,omitempty" xml:"e_code,omitempty"`
    // error_msg
    EMsg   string `json:"e_msg,omitempty" xml:"e_msg,omitempty"`
}
