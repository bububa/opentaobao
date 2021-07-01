package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

/* AlibabaSecurityJaqCaptchaSend
聚安全安全验证发起接口
alibaba.security.jaq.captcha.send

聚安全安全验证发起 */
func AlibabaSecurityJaqCaptchaSend(clt *core.SDKClient, req *security.AlibabaSecurityJaqCaptchaSendAPIRequest, session string) (*security.AlibabaSecurityJaqCaptchaSendAPIResponse, error) {
	var resp security.AlibabaSecurityJaqCaptchaSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
