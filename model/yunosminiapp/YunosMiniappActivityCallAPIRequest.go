package yunosminiapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosMiniappActivityCallAPIRequest 调用活动接口 API请求
// yunos.miniapp.activity.call
//
// 用于小程序调用活动接口
type YunosMiniappActivityCallAPIRequest struct {
	model.Params
	// 活动id
	_activityId string
	// 设备id
	_deviceId string
	// 请求选项
	_options *Options
}

// NewYunosMiniappActivityCallRequest 初始化YunosMiniappActivityCallAPIRequest对象
func NewYunosMiniappActivityCallRequest() *YunosMiniappActivityCallAPIRequest {
	return &YunosMiniappActivityCallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosMiniappActivityCallAPIRequest) GetApiMethodName() string {
	return "yunos.miniapp.activity.call"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosMiniappActivityCallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosMiniappActivityCallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *YunosMiniappActivityCallAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r YunosMiniappActivityCallAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *YunosMiniappActivityCallAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r YunosMiniappActivityCallAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetOptions is Options Setter
// 请求选项
func (r *YunosMiniappActivityCallAPIRequest) SetOptions(_options *Options) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r YunosMiniappActivityCallAPIRequest) GetOptions() *Options {
	return r._options
}
