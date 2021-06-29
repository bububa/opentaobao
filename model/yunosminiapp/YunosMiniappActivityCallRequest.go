package yunosminiapp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调用活动接口 APIRequest
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

func NewYunosMiniappActivityCallRequest() *YunosMiniappActivityCallRequest{
    return &YunosMiniappActivityCallRequest{
        Params: model.NewParams(),
    }
}

func (r YunosMiniappActivityCallRequest) GetApiMethodName() string {
    return "yunos.miniapp.activity.call"
}

func (r YunosMiniappActivityCallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosMiniappActivityCallRequest) SetOptions(options *Options) error {
    r.options = options
    r.Set("options", options)
    return nil
}

func (r YunosMiniappActivityCallRequest) GetOptions() *Options {
    return r.options
}

func (r *YunosMiniappActivityCallRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r YunosMiniappActivityCallRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *YunosMiniappActivityCallRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r YunosMiniappActivityCallRequest) GetActivityId() string {
    return r.activityId
}

