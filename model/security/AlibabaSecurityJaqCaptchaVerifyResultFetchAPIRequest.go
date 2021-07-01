package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest
聚安全安全验证检查结果获取接口 API请求
alibaba.security.jaq.captcha.verify.result.fetch

获取二次验证的结果 */
type AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest struct {
	model.Params
	// 二次验证获取验证检查结果所需的seesionId
	_sessionId string
}

// New
