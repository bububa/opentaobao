package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqcaptchaverifyresultfetch 聚安全安全验证检查结果获取接口
// alibaba.security.jaq.captcha.verify.result.fetch
//
// 获取二次验证的结果
func Alibabasecurityjaqcaptchaverifyresultfetch(clt *core.SDKClient, req *security.AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest, session string) (*security.AlibabasecurityjaqcaptchaverifyresultfetchAPIResponse, error) {
	var resp security.AlibabasecurityjaqcaptchaverifyresultfetchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
