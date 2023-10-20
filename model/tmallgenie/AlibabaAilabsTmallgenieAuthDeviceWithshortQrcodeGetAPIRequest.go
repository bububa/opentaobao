package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest 根据安全简码查询二维码详细信息 API请求
// alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get
//
// 根据安全简码查询二维码详细信息
type AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest struct {
	model.Params
	// 产品ID
	_clientId string
	// 授权码
	_authCode string
}

// NewAlibabaailabstmallgenieauthdevicewithshortqrcodegetRequest 初始化AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicewithshortqrcodegetRequest() *AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest {
	return &AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 产品ID
func (r *AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetAuthCode is AuthCode Setter
// 授权码
func (r *AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest) SetAuthCode(_authCode string) error {
	r._authCode = _authCode
	r.Set("auth_code", _authCode)
	return nil
}

// GetAuthCode AuthCode Getter
func (r AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest) GetAuthCode() string {
	return r._authCode
}
