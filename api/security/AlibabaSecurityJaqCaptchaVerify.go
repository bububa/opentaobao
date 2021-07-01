package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

/* AlibabaSecurityJaqCaptchaVerify
聚安全安全验证检查接口
alibaba.security.jaq.captcha.verify

聚安全安全验证检查 */
func AlibabaSecurityJaqCaptchaVerify(clt *core.SDKClient, req *security.AlibabaSecurityJaqCaptchaVerifyAPIRequest, session string) (*security.AlibabaSecurityJaqCaptchaVerifyAPIResponse, error) {
	var resp security.AlibabaSecurityJaqCaptchaVerifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
