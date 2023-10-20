package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotdevicelistgetAPIRequest 获取iot设备列表 API请求
// alibaba.ailabs.iot.device.list.get
//
// 通过此接口获取用户名下的iot设备列表
type AlibabaailabsiotdevicelistgetAPIRequest struct {
	model.Params
	// 用户open id
	_userOpenId string
	// client id
	_clientId string
}

// NewAlibabaailabsiotdevicelistgetRequest 初始化AlibabaailabsiotdevicelistgetAPIRequest对象
func NewAlibabaailabsiotdevicelistgetRequest() *AlibabaailabsiotdevicelistgetAPIRequest {
	return &AlibabaailabsiotdevicelistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotdevicelistgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotdevicelistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotdevicelistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserOpenId is UserOpenId Setter
// 用户open id
func (r *AlibabaailabsiotdevicelistgetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaailabsiotdevicelistgetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetClientId is ClientId Setter
// client id
func (r *AlibabaailabsiotdevicelistgetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabsiotdevicelistgetAPIRequest) GetClientId() string {
	return r._clientId
}
