package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicegetAPIRequest 获取设备详情 API请求
// alibaba.ailabs.tmallgenie.auth.device.get
//
// 通过此接口获取设备详情
type AlibabaailabstmallgenieauthdevicegetAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
	// 设备uuid
	_uuid string
}

// NewAlibabaailabstmallgenieauthdevicegetRequest 初始化AlibabaailabstmallgenieauthdevicegetAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicegetRequest() *AlibabaailabstmallgenieauthdevicegetAPIRequest {
	return &AlibabaailabstmallgenieauthdevicegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicegetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 客户id
func (r *AlibabaailabstmallgenieauthdevicegetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthdevicegetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetUserOpenId is UserOpenId Setter
// 用户开放id
func (r *AlibabaailabstmallgenieauthdevicegetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaailabstmallgenieauthdevicegetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabaailabstmallgenieauthdevicegetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaailabstmallgenieauthdevicegetAPIRequest) GetUuid() string {
	return r._uuid
}
