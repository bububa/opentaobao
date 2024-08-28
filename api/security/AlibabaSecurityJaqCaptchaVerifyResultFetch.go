package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqCaptchaVerifyResultFetch 聚安全安全验证检查结果获取接口
// alibaba.security.jaq.captcha.verify.result.fetch
//
// 获取二次验证的结果
func AlibabaSecurityJaqCaptchaVerifyResultFetch(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest, resp *security.AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
