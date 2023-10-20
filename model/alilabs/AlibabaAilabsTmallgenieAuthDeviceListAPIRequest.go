package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicelistAPIRequest 获取用户设备列表 API请求
// alibaba.ailabs.tmallgenie.auth.device.list
//
// 通过此接口获取用户绑定的设备信息列表
type AlibabaailabstmallgenieauthdevicelistAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
}

// NewAlibabaailabstmallgenieauthdevicelistRequest 初始化AlibabaailabstmallgenieauthdevicelistAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicelistRequest() *AlibabaailabstmallgenieauthdevicelistAPIRequest {
	return &AlibabaailabstmallgenieauthdevicelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicelistAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 客户id
func (r *AlibabaailabstmallgenieauthdevicelistAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthdevicelistAPIRequest) GetClientId() string {
	return r._clientId
}

// SetUserOpenId is UserOpenId Setter
// 用户开放id
func (r *AlibabaailabstmallgenieauthdevicelistAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaailabstmallgenieauthdevicelistAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}
