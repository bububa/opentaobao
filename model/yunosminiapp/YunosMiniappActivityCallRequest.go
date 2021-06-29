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
type YunosMiniappActivityCallRequest struct {
    model.Params
    // 请求选项
    options   *Options
    // 设备id
    deviceId   string
    // 活动id
    activityId   string
}

// 初始化YunosMiniappActivityCallRequest对象
func NewYunosMiniappActivityCallRequest() *YunosMiniappActivityCallRequest{
    return &YunosMiniappActivityCallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosMiniappActivityCallRequest) GetApiMethodName() string {
    return "yunos.miniapp.activity.call"
}

// IRequest interface 方法, 获取API参数
func (r YunosMiniappActivityCallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Options Setter
// 请求选项
func (r *YunosMiniappActivityCallRequest) SetOptions(options *Options) error {
    r.options = options
    r.Set("options", options)
    return nil
}

// Options Getter
func (r YunosMiniappActivityCallRequest) GetOptions() *Options {
    return r.options
}
// DeviceId Setter
// 设备id
func (r *YunosMiniappActivityCallRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r YunosMiniappActivityCallRequest) GetDeviceId() string {
    return r.deviceId
}
// ActivityId Setter
// 活动id
func (r *YunosMiniappActivityCallRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r YunosMiniappActivityCallRequest) GetActivityId() string {
    return r.activityId
}
