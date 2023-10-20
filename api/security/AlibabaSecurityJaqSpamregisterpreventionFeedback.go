package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqSpamregisterpreventionFeedback 保护结果反馈
// alibaba.security.jaq.spamregisterprevention.feedback
//
// 用户通过这个接口对垃圾注册防控结果进行反馈
func AlibabaSecurityJaqSpamregisterpreventionFeedback(clt *core.SDKClient, req *security.AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest, resp *security.AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
