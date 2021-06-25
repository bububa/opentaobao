package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
保护结果反馈 APIResponse
alibaba.security.jaq.spamregisterprevention.feedback

用户通过这个接口对垃圾注册防控结果进行反馈
*/
type AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqSpamregisterpreventionFeedbackResponse `json:"alibaba_security_jaq_spamregisterprevention_feedback_response,omitempty"`
}

type AlibabaSecurityJaqSpamregisterpreventionFeedbackResponse struct {

    // feedBack返回结果
    FeedBackResult   *JaqFeedBackResult `json:"feed_back_result,omitempty"`

}
