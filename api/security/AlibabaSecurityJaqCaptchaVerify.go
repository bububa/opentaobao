package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqCaptchaVerify 聚安全安全验证检查接口
// alibaba.security.jaq.captcha.verify
//
// 聚安全安全验证检查
func AlibabaSecurityJaqCaptchaVerify(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqCaptchaVerifyAPIRequest, resp *security.AlibabaSecurityJaqCaptchaVerifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
