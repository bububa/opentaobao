package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest
设备状态查询 API请求
alibaba.ailabs.tmallgenie.auth.device.status.get

提供给天猫精灵定制机厂商 查询设备在线状态值 */
type AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest struct {
	model.Params
	// 给产品分配的唯一ID（22位）
	_clientId string
	// 精灵用户的唯一外部ID
	_userOpenId string
	// 精灵设备的唯一ID
	_uuid string
}

// NewAlibabaAilabsTmallgenieAuthDeviceStatusGetRequest 初始化AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceStatusGetRequest() *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ClientId Setter
// 给产品分配的唯一ID（22位）
func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// Get ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetClientId() string {
	return r._clientId
}

// Set is UserOpenId Setter
// 精灵用户的唯一外部ID
func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// Get UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// Set is Uuid Setter
// 精灵设备的唯一ID
func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetUuid() string {
	return r._uuid
}
