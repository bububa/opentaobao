package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudSubmitAPIRequest 实人认证云服务提交接口 API请求
// alibaba.security.jaq.rp.cloud.submit
//
// 聚安全实人认证提交认证接口
type AlibabaSecurityJaqRpCloudSubmitAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
}

// NewAlibabaSecurityJaqRpCloudSubmitRequest 初始化AlibabaSecurityJaqRpCloudSubmitAPIRequest对象
func NewAlibabaSecurityJaqRpCloudSubmitRequest() *AlibabaSecurityJaqRpCloudSubmitAPIRequest {
	return &AlibabaSecurityJaqRpCloudSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVerifyToken is VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpCloudSubmitAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudSubmitAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}
