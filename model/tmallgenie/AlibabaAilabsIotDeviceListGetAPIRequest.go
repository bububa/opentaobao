package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceListGetAPIRequest 获取iot设备列表 API请求
// alibaba.ailabs.iot.device.list.get
//
// 通过此接口获取用户名下的iot设备列表
type AlibabaAilabsIotDeviceListGetAPIRequest struct {
	model.Params
	// 用户open id
	_userOpenId string
	// client id
	_clientId string
}

// NewAlibabaAilabsIotDeviceListGetRequest 初始化AlibabaAilabsIotDeviceListGetAPIRequest对象
func NewAlibabaAilabsIotDeviceListGetRequest() *AlibabaAilabsIotDeviceListGetAPIRequest {
	return &AlibabaAilabsIotDeviceListGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserOpenId Setter
// 用户open id
func (r *AlibabaAilabsIotDeviceListGetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// Get UserOpenId Getter
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// Set is ClientId Setter
// client id
func (r *AlibabaAilabsIotDeviceListGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// Get ClientId Getter
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetClientId() string {
	return r._clientId
}
