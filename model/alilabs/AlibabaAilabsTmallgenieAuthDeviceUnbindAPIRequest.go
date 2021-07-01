package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest
解绑设备 API请求
alibaba.ailabs.tmallgenie.auth.device.unbind

通过此接口解绑天猫精灵设备 */
type AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
	// 设备uuid
	_uuid string
}

// NewAlibabaAilabsTmallgenieAuthDeviceUnbindRequest 初始化AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceUnbindRequest() *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// Get ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetClientId() string {
	return r._clientId
}

// Set is UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// Get UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// Set is Uuid Setter
// 设备uuid
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetUuid() string {
	return r._uuid
}
