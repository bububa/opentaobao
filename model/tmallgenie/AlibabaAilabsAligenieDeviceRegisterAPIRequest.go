package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligeniedeviceregisterAPIRequest 天猫精灵开放平台获取设备秘钥 API请求
// alibaba.ailabs.aligenie.device.register
//
// 向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥
type AlibabaailabsaligeniedeviceregisterAPIRequest struct {
	model.Params
	// mac区段脚本
	_macSections string
	// 设备id
	_deviceId int64
}

// NewAlibabaailabsaligeniedeviceregisterRequest 初始化AlibabaailabsaligeniedeviceregisterAPIRequest对象
func NewAlibabaailabsaligeniedeviceregisterRequest() *AlibabaailabsaligeniedeviceregisterAPIRequest {
	return &AlibabaailabsaligeniedeviceregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligeniedeviceregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.device.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligeniedeviceregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligeniedeviceregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMacSections is MacSections Setter
// mac区段脚本
func (r *AlibabaailabsaligeniedeviceregisterAPIRequest) SetMacSections(_macSections string) error {
	r._macSections = _macSections
	r.Set("mac_sections", _macSections)
	return nil
}

// GetMacSections MacSections Getter
func (r AlibabaailabsaligeniedeviceregisterAPIRequest) GetMacSections() string {
	return r._macSections
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *AlibabaailabsaligeniedeviceregisterAPIRequest) SetDeviceId(_deviceId int64) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaailabsaligeniedeviceregisterAPIRequest) GetDeviceId() int64 {
	return r._deviceId
}
