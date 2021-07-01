package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest
获取设备详情 API请求
alibaba.ailabs.tmallgenie.auth.device.get

通过此接口获取设备详情 */
type AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
	// 设备uuid
	_uuid string
}

// NewAlibabaAilabsTmallgenieAuthDeviceGetRequest 初始化AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceGetRequest() *AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// Get ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest) GetClientId() string {
	return r._clientId
}

// Set is UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// Get UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// Set is Uuid Setter
// 设备uuid
func (r *AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest) GetUuid() string {
	return r._uuid
}
