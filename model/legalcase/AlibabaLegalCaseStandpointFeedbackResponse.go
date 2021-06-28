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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_legal_case_standpoint_feedback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"