package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthgettokenAPIRequest 设备授权 API请求
// alibaba.ailabs.tmallgenie.auth.gettoken
//
// 获取设备授权码
type AlibabaailabstmallgenieauthgettokenAPIRequest struct {
	model.Params
	// clientId
	_clientId string
	// 授权码
	_authCode string
	// 授权类型，只支持authorization_code
	_grantType string
}

// NewAlibabaailabstmallgenieauthgettokenRequest 初始化AlibabaailabstmallgenieauthgettokenAPIRequest对象
func NewAlibabaailabstmallgenieauthgettokenRequest() *AlibabaailabstmallgenieauthgettokenAPIRequest {
	return &AlibabaailabstmallgenieauthgettokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthgettokenAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.gettoken"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthgettokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthgettokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// clientId
func (r *AlibabaailabstmallgenieauthgettokenAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthgettokenAPIRequest) GetClientId() string {
	return r._clientId
}

// SetAuthCode is AuthCode Setter
// 授权码
func (r *AlibabaailabstmallgenieauthgettokenAPIRequest) SetAuthCode(_authCode string) error {
	r._authCode = _authCode
	r.Set("auth_code", _authCode)
	return nil
}

// GetAuthCode AuthCode Getter
func (r AlibabaailabstmallgenieauthgettokenAPIRequest) GetAuthCode() string {
	return r._authCode
}

// SetGrantType is GrantType Setter
// 授权类型，只支持authorization_code
func (r *AlibabaailabstmallgenieauthgettokenAPIRequest) SetGrantType(_grantType string) error {
	r._grantType = _grantType
	r.Set("grant_type", _grantType)
	return nil
}

// GetGrantType GrantType Getter
func (r AlibabaailabstmallgenieauthgettokenAPIRequest) GetGrantType() string {
	return r._grantType
}
