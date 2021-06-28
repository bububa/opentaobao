package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息通知 APIResponse
alibaba.legal.case.common.notice

同步通知给诉讼系统
*/
type AlibabaLegalCaseCommonNoticeAPIResponse struct {
    model.CommonResponse
    AlibabaLegalCaseCommonNoticeResponse
}

type AlibabaLegalCaseCommonNoticeResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_case_common_notice_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // error
    
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`

    
    // content
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`

    
    // msg
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
