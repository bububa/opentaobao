package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqcaptchaverify 聚安全安全验证检查接口
// alibaba.security.jaq.captcha.verify
//
// 聚安全安全验证检查
func Alibabasecurityjaqcaptchaverify(clt *core.SDKClient, req *security.AlibabasecurityjaqcaptchaverifyAPIRequest, session string) (*security.AlibabasecurityjaqcaptchaverifyAPIResponse, error) {
	var resp security.AlibabasecurityjaqcaptchaverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
