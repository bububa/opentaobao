package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqSpamregisterpreventionFeedback 保护结果反馈
// alibaba.security.jaq.spamregisterprevention.feedback
//
// 用户通过这个接口对垃圾注册防控结果进行反馈
func AlibabaSecurityJaqSpamregisterpreventionFeedback(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest, resp *security.AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
