package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增反馈口径 APIResponse
alibaba.legal.case.standpoint.savestandpoint

新增反馈口径 ,从外部接受反馈的口径
*/
type AlibabaLegalCaseStandpointSavestandpointAPIResponse struct {
    model.CommonResponse
    Response *AlibabaLegalCaseStandpointSavestandpointResponse `json:"alibaba_legal_case_standpoint_savestandpoint_response,omitempty"`
}

type AlibabaLegalCaseStandpointSavestandpointResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误编码
    Errcode   string `json:"errcode,omitempty"`

    // 反馈的新增口径id
    Content   int64 `json:"content,omitempty"`

    // 错误描述
    Errmsg   string `json:"errmsg,omitempty"`

}
