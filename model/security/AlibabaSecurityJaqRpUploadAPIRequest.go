package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpUploadAPIRequest 聚安全实人认证上传认证信息 API请求
// alibaba.security.jaq.rp.upload
//
// 聚安全实人认证上传认证信息
type AlibabaSecurityJaqRpUploadAPIRequest struct {
	model.Params
	// 认证会话token
	_verifyToken string
	// 此次需要上传的认证信息的列表
	_elements []Element
}

// NewAlibabaSecurityJaqRpUploadRequest 初始化AlibabaSecurityJaqRpUploadAPIRequest对象
func NewAlibabaSecurityJaqRpUploadRequest() *AlibabaSecurityJaqRpUploadAPIRequest {
	return &AlibabaSecurityJaqRpUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVerifyToken is VerifyToken Setter
// 认证会话token
func (r *AlibabaSecurityJaqRpUploadAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpUploadAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetElements is Elements Setter
// 此次需要上传的认证信息的列表
func (r *AlibabaSecurityJaqRpUploadAPIRequest) SetElements(_elements []Element) error {
	r._elements = _elements
	r.Set("elements", _elements)
	return nil
}

// GetElements Elements Getter
func (r AlibabaSecurityJaqRpUploadAPIRequest) GetElements() []Element {
	return r._elements
}
