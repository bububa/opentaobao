package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增或更新 反馈口径(采纳口径/不采纳口径) APIResponse
alibaba.legal.case.standpoint.feedback

新增或更新 反馈口径(采纳口径/不采纳口径)
*/
type AlibabaLegalCaseStandpointFeedbackAPIResponse struct {
    model.CommonResponse
    AlibabaLegalCaseStandpointFeedbackResponse
}

type AlibabaLegalCaseStandpointFeedbackResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_case_standpoint_feedback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误编码
    
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`

    
    // 成功失败 true/false
    
    Content   bool `json:"content,omitempty" xml:"content,omitempty"`

    
    // 错误信息描述
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
