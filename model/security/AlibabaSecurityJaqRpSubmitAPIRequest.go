package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpSubmitAPIRequest 聚安全实人认证提交认证接口 API请求
// alibaba.security.jaq.rp.submit
//
// 聚安全实人认证提交认证接口
type AlibabaSecurityJaqRpSubmitAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
}

// NewAlibabaSecurityJaqRpSubmitRequest 初始化AlibabaSecurityJaqRpSubmitAPIRequest对象
func NewAlibabaSecurityJaqRpSubmitRequest() *AlibabaSecurityJaqRpSubmitAPIRequest {
	return &AlibabaSecurityJaqRpSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpSubmitAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpSubmitAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}
