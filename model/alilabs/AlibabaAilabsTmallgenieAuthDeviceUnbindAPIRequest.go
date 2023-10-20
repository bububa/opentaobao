package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdeviceunbindAPIRequest 解绑设备 API请求
// alibaba.ailabs.tmallgenie.auth.device.unbind
//
// 通过此接口解绑天猫精灵设备
type AlibabaailabstmallgenieauthdeviceunbindAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
	// 设备uuid
	_uuid string
}

// NewAlibabaailabstmallgenieauthdeviceunbindRequest 初始化AlibabaailabstmallgenieauthdeviceunbindAPIRequest对象
func NewAlibabaailabstmallgenieauthdeviceunbindRequest() *AlibabaailabstmallgenieauthdeviceunbindAPIRequest {
	return &AlibabaailabstmallgenieauthdeviceunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdeviceunbindAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdeviceunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdeviceunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 客户id
func (r *AlibabaailabstmallgenieauthdeviceunbindAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthdeviceunbindAPIRequest) GetClientId() string {
	return r._clientId
}

// SetUserOpenId is UserOpenId Setter
// 用户开放id
func (r *AlibabaailabstmallgenieauthdeviceunbindAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaailabstmallgenieauthdeviceunbindAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabaailabstmallgenieauthdeviceunbindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaailabstmallgenieauthdeviceunbindAPIRequest) GetUuid() string {
	return r._uuid
}
