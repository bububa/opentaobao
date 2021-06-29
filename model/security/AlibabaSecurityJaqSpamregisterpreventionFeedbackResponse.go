package security

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保护结果反馈 APIResponse
alibaba.security.jaq.spamregisterprevention.feedback

用户通过这个接口对垃圾注册防控结果进行反馈
*/
type AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse struct {
    model.CommonResponse
    AlibabaSecurityJaqSpamregisterpreventionFeedbackResponse
}

type AlibabaSecurityJaqSpamregisterpreventionFeedbackResponse struct {
    XMLName xml.Name `xml:"alibaba_security_jaq_spamregisterprevention_feedback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // feedBack返回结果
    
    FeedBackResult   *JaqFeedBackResult `json:"feed_back_result,omitempty" xml:"feed_back_result,omitempty"`

    
}
