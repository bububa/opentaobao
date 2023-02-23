package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest 获取mqtt设备信息 API请求
// taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken
//
// 智能硬件设备动态注册和获取mqtt设备信息
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest struct {
	model.Params
	// 设备的唯一标识码，比如网卡的MAC地址
	_deviceId string
	// 访问令牌
	_applyAccessToken string
	// 业务编码，具体编码请联系杜尘
	_businessCode int64
	// 设备类型编码，具体编码请联系杜尘
	_deviceType int64
	// 环境编码，0日常，1预发，2线上
	_enviroCode int64
}

// NewTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest 初始化TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest对象
func NewTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest() *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest {
	return &TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备的唯一标识码，比如网卡的MAC地址
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetApplyAccessToken is ApplyAccessToken Setter
// 访问令牌
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) SetApplyAccessToken(_applyAccessToken string) error {
	r._applyAccessToken = _applyAccessToken
	r.Set("apply_access_token", _applyAccessToken)
	return nil
}

// GetApplyAccessToken ApplyAccessToken Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) GetApplyAccessToken() string {
	return r._applyAccessToken
}

// SetBusinessCode is BusinessCode Setter
// 业务编码，具体编码请联系杜尘
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) SetBusinessCode(_businessCode int64) error {
	r._businessCode = _businessCode
	r.Set("business_code", _businessCode)
	return nil
}

// GetBusinessCode BusinessCode Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) GetBusinessCode() int64 {
	return r._businessCode
}

// SetDeviceType is DeviceType Setter
// 设备类型编码，具体编码请联系杜尘
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// SetEnviroCode is EnviroCode Setter
// 环境编码，0日常，1预发，2线上
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) SetEnviroCode(_enviroCode int64) error {
	r._enviroCode = _enviroCode
	r.Set("enviro_code", _enviroCode)
	return nil
}

// GetEnviroCode EnviroCode Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIRequest) GetEnviroCode() int64 {
	return r._enviroCode
}
