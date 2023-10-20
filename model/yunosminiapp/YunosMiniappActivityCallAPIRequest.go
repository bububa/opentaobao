package yunosminiapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosminiappactivitycallAPIRequest 调用活动接口 API请求
// yunos.miniapp.activity.call
//
// 用于小程序调用活动接口
type YunosminiappactivitycallAPIRequest struct {
	model.Params
	// 活动id
	_activityId string
	// 设备id
	_deviceId string
	// 请求选项
	_options *Options
}

// NewYunosminiappactivitycallRequest 初始化YunosminiappactivitycallAPIRequest对象
func NewYunosminiappactivitycallRequest() *YunosminiappactivitycallAPIRequest {
	return &YunosminiappactivitycallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosminiappactivitycallAPIRequest) GetApiMethodName() string {
	return "yunos.miniapp.activity.call"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosminiappactivitycallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosminiappactivitycallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *YunosminiappactivitycallAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r YunosminiappactivitycallAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *YunosminiappactivitycallAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r YunosminiappactivitycallAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetOptions is Options Setter
// 请求选项
func (r *YunosminiappactivitycallAPIRequest) SetOptions(_options *Options) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r YunosminiappactivitycallAPIRequest) GetOptions() *Options {
	return r._options
}
