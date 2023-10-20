package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthswitchuserAPIRequest 切换用户 API请求
// alibaba.ailabs.tmallgenie.auth.switchuser
//
// 设备切换授权用户
type AlibabaailabstmallgenieauthswitchuserAPIRequest struct {
	model.Params
	// client_id
	_clientId string
	// 目标用户openId
	_newUserOpenId string
	// 当前拥有设备权限的用户openId
	_oldUserOpenId string
	// 设备uuid
	_uuid string
}

// NewAlibabaailabstmallgenieauthswitchuserRequest 初始化AlibabaailabstmallgenieauthswitchuserAPIRequest对象
func NewAlibabaailabstmallgenieauthswitchuserRequest() *AlibabaailabstmallgenieauthswitchuserAPIRequest {
	return &AlibabaailabstmallgenieauthswitchuserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthswitchuserAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.switchuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthswitchuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthswitchuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// client_id
func (r *AlibabaailabstmallgenieauthswitchuserAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthswitchuserAPIRequest) GetClientId() string {
	return r._clientId
}

// SetNewUserOpenId is NewUserOpenId Setter
// 目标用户openId
func (r *AlibabaailabstmallgenieauthswitchuserAPIRequest) SetNewUserOpenId(_newUserOpenId string) error {
	r._newUserOpenId = _newUserOpenId
	r.Set("new_user_open_id", _newUserOpenId)
	return nil
}

// GetNewUserOpenId NewUserOpenId Getter
func (r AlibabaailabstmallgenieauthswitchuserAPIRequest) GetNewUserOpenId() string {
	return r._newUserOpenId
}

// SetOldUserOpenId is OldUserOpenId Setter
// 当前拥有设备权限的用户openId
func (r *AlibabaailabstmallgenieauthswitchuserAPIRequest) SetOldUserOpenId(_oldUserOpenId string) error {
	r._oldUserOpenId = _oldUserOpenId
	r.Set("old_user_open_id", _oldUserOpenId)
	return nil
}

// GetOldUserOpenId OldUserOpenId Getter
func (r AlibabaailabstmallgenieauthswitchuserAPIRequest) GetOldUserOpenId() string {
	return r._oldUserOpenId
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabaailabstmallgenieauthswitchuserAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaailabstmallgenieauthswitchuserAPIRequest) GetUuid() string {
	return r._uuid
}
