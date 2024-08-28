package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqCaptchaSend 聚安全安全验证发起接口
// alibaba.security.jaq.captcha.send
//
// 聚安全安全验证发起
func AlibabaSecurityJaqCaptchaSend(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqCaptchaSendAPIRequest, resp *security.AlibabaSecurityJaqCaptchaSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
