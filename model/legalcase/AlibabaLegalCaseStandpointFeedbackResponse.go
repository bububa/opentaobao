package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增或更新 反馈口径(采纳口径/不采纳口径) APIResponse
alibaba.legal.case.standpoint.feedback

新增或更新 反馈口径(采纳口径/不采纳口径)
*/
type AlibabaLegalCaseStandpointFeedbackAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLegalCaseStandpointFeedbackResponse `json:"alibaba_legal_case_standpoint_feedback_response,omitempty"` 
    AlibabaLegalCaseStandpointFeedbackResponse
}

/* model for simplify = false
type AlibabaLegalCaseStandpointFeedbackResponse struct {

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误编码
    
    Errcode   string `json:"errcode,omitempty"`
    

    // 成功失败 true/false
    
    Content   bool `json:"content,omitempty"`
    

    // 错误信息描述
    
    Errmsg   string `json:"errmsg,omitempty"`
    

}
*/

type AlibabaLegalCaseStandpointFeedbackResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误编码
    Errcode   string `json:"errcode,omitempty"`

    // 成功失败 true/false
    Content   bool `json:"content,omitempty"`

    // 错误信息描述
    Errmsg   string `json:"errmsg,omitempty"`

}
