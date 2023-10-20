package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest 聚安全安全验证检查结果获取接口 API请求
// alibaba.security.jaq.captcha.verify.result.fetch
//
// 获取二次验证的结果
type AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest struct {
	model.Params
	// 二次验证获取验证检查结果所需的seesionId
	_sessionId string
}

// NewAlibabasecurityjaqcaptchaverifyresultfetchRequest 初始化AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest对象
func NewAlibabasecurityjaqcaptchaverifyresultfetchRequest() *AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest {
	return &AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.captcha.verify.result.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSessionId is SessionId Setter
// 二次验证获取验证检查结果所需的seesionId
func (r *AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// GetSessionId SessionId Getter
func (r AlibabasecurityjaqcaptchaverifyresultfetchAPIRequest) GetSessionId() string {
	return r._sessionId
}
