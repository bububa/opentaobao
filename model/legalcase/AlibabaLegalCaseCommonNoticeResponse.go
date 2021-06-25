package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息通知 APIResponse
alibaba.legal.case.common.notice

同步通知给诉讼系统
*/
type AlibabaLegalCaseCommonNoticeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaLegalCaseCommonNoticeResponse `json:"alibaba_legal_case_common_notice_response,omitempty"`
}

type AlibabaLegalCaseCommonNoticeResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // error
    Errcode   string `json:"errcode,omitempty"`

    // content
    Content   string `json:"content,omitempty"`

    // msg
    Errmsg   string `json:"errmsg,omitempty"`

}
