package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicestatusgetAPIRequest 设备状态查询 API请求
// alibaba.ailabs.tmallgenie.auth.device.status.get
//
// 提供给天猫精灵定制机厂商 查询设备在线状态值
type AlibabaailabstmallgenieauthdevicestatusgetAPIRequest struct {
	model.Params
	// 给产品分配的唯一ID（22位）
	_clientId string
	// 精灵用户的唯一外部ID
	_userOpenId string
	// 精灵设备的唯一ID
	_uuid string
}

// NewAlibabaailabstmallgenieauthdevicestatusgetRequest 初始化AlibabaailabstmallgenieauthdevicestatusgetAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicestatusgetRequest() *AlibabaailabstmallgenieauthdevicestatusgetAPIRequest {
	return &AlibabaailabstmallgenieauthdevicestatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 给产品分配的唯一ID（22位）
func (r *AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetUserOpenId is UserOpenId Setter
// 精灵用户的唯一外部ID
func (r *AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetUuid is Uuid Setter
// 精灵设备的唯一ID
func (r *AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaailabstmallgenieauthdevicestatusgetAPIRequest) GetUuid() string {
	return r._uuid
}
