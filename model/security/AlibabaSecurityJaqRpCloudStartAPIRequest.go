package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudStartAPIRequest 实人认证云开始认证 API请求
// alibaba.security.jaq.rp.cloud.start
//
// 聚安全实人认证开始
type AlibabaSecurityJaqRpCloudStartAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 客户端信息，如果是服务端接入，里面的参数可为空
	_clientInfo *RpClientInfo
	// 扩展信息
	_extraData string
}

// NewAlibabaSecurityJaqRpCloudStartRequest 初始化AlibabaSecurityJaqRpCloudStartAPIRequest对象
func NewAlibabaSecurityJaqRpCloudStartRequest() *AlibabaSecurityJaqRpCloudStartAPIRequest {
	return &AlibabaSecurityJaqRpCloudStartAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.start"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpCloudStartAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetClientInfo is ClientInfo Setter
// 客户端信息，如果是服务端接入，里面的参数可为空
func (r *AlibabaSecurityJaqRpCloudStartAPIRequest) SetClientInfo(_clientInfo *RpClientInfo) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetClientInfo() *RpClientInfo {
	return r._clientInfo
}

// SetExtraData is ExtraData Setter
// 扩展信息
func (r *AlibabaSecurityJaqRpCloudStartAPIRequest) SetExtraData(_extraData string) error {
	r._extraData = _extraData
	r.Set("extra_data", _extraData)
	return nil
}

// GetExtraData ExtraData Getter
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetExtraData() string {
	return r._extraData
}
