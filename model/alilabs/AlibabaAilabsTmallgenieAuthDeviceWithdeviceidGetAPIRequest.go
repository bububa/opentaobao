package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest 根据三方ID查询设备注册激活信息 API请求
// alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get
//
// 根据三方ID查询设备注册激活信息
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest struct {
	model.Params
	// 设备产品ID
	_clientId string
	// mac地址
	_mac string
}

// NewAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest 初始化AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest() *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ClientId Setter
// 设备产品ID
func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// Get ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetClientId() string {
	return r._clientId
}

// Set is Mac Setter
// mac地址
func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// Get Mac Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetMac() string {
	return r._mac
}
