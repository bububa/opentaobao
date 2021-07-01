package yunosminiapp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调用活动接口 API请求
yunos.miniapp.activity.call

用于小程序调用活动接口
*/
type YunosMiniappActivityCallAPIRequest struct {
    model.Params
    // 请求选项
    _options   *Options
    // 设备id
    _deviceId   string
    // 活动id
    _activityId   string
}

// 初始化YunosMiniappActivityCallAPIRequest对象
func NewYunosMiniappActivityCallRequest() *YunosMiniappActivityCallAPIRequest{
    return &YunosMiniappActivityCallAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosMiniappActivityCallAPIRequest) GetApiMethodName() string {
    return "yunos.miniapp.activity.call"
}

// IRequest interface 方法, 获取API参数
func (r YunosMiniappActivityCallAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Options Setter
// 请求选项
func (r *YunosMiniappActivityCallAPIRequest) SetOptions(_options *Options) error {
    r._options = _options
    r.Set("options", _options)
    return nil
}

// Options Getter
func (r YunosMiniappActivityCallAPIRequest) GetOptions() *Options {
    return r._options
}
// DeviceId Setter
// 设备id
func (r *YunosMiniappActivityCallAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r YunosMiniappActivityCallAPIRequest) GetDeviceId() string {
    return r._deviceId
}
// ActivityId Setter
// 活动id
func (r *YunosMiniappActivityCallAPIRequest) SetActivityId(_activityId string) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r YunosMiniappActivityCallAPIRequest) GetActivityId() string {
    return r._activityId
}
