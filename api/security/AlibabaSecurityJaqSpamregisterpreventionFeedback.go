package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
保护结果反馈 
alibaba.security.jaq.spamregisterprevention.feedback

用户通过这个接口对垃圾注册防控结果进行反馈
*/
func AlibabaSecurityJaqSpamregisterpreventionFeedback(clt *core.SDKClient, req *security.AlibabaSecurityJaqSpamregisterpreventionFeedbackRequest, session string) (*security.AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse, error) {
    var resp security.AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
