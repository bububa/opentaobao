package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqcaptchasend 聚安全安全验证发起接口
// alibaba.security.jaq.captcha.send
//
// 聚安全安全验证发起
func Alibabasecurityjaqcaptchasend(clt *core.SDKClient, req *security.AlibabasecurityjaqcaptchasendAPIRequest, session string) (*security.AlibabasecurityjaqcaptchasendAPIResponse, error) {
	var resp security.AlibabasecurityjaqcaptchasendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
