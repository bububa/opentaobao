package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest 聚安全安全验证检查结果获取接口 API请求
// alibaba.security.jaq.captcha.verify.result.fetch
//
// 获取二次验证的结果
type AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest struct {
	model.Params
	// 二次验证获取验证检查结果所需的seesionId
	_sessionId string
}

// NewAlibabaSecurityJaqCaptchaVerifyResultFetchRequest 初始化AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest对象
func NewAlibabaSecurityJaqCaptchaVerifyResultFetchRequest() *AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest {
	return &AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.captcha.verify.result.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSessionId is SessionId Setter
// 二次验证获取验证检查结果所需的seesionId
func (r *AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// GetSessionId SessionId Getter
func (r AlibabaSecurityJaqCaptchaVerifyResultFetchAPIRequest) GetSessionId() string {
	return r._sessionId
}
